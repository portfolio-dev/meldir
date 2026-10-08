package http

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strings"

	"meldir-backend/internal/delivery/http/handler"
	"meldir-backend/internal/delivery/http/middleware"
	"meldir-backend/internal/domain"
)

func NewRouter(
	healthHandler *handler.HealthHandler,
	authHandler *handler.AuthHandler,
	userHandler *handler.UserHandler,
	ticketHandler *handler.TicketHandler,
	timesheetHandler *handler.TimesheetHandler,
	invoiceHandler *handler.InvoiceHandler,
	accountingHandler *handler.AccountingHandler,
	leadHandler *handler.LeadHandler,
	authMiddleware *middleware.AuthMiddleware,
) http.Handler {
	mux := http.NewServeMux()

	// 1. Health check & Ping endpoints
	mux.HandleFunc("/api/health", healthHandler.CheckHealth)
	mux.HandleFunc("/api/v1/ping", healthHandler.Ping)

	// 2. Auth Endpoints
	mux.HandleFunc("/api/v1/auth/login", authHandler.Login)
	mux.HandleFunc("/api/v1/user/profile", authMiddleware.Authenticate(authHandler.GetProfile))
	mux.HandleFunc("/api/v1/auth/logout", authMiddleware.Authenticate(authHandler.Logout))

	// 3. User Management CRUD (Khusus Direktur Utama & Admin)
	adminOnly := authMiddleware.RequireRoles(domain.RoleDirektur, domain.RoleAdmin)
	mux.HandleFunc("/api/v1/users", adminOnly(userHandler.ListUsers))
	mux.HandleFunc("/api/v1/users/create", adminOnly(userHandler.CreateUser))
	mux.HandleFunc("/api/v1/users/update", adminOnly(userHandler.UpdateUser))
	mux.HandleFunc("/api/v1/users/delete", adminOnly(userHandler.DeleteUser))

	// 4. SLA Tickets Endpoints (Otentikasi Pengguna)
	mux.HandleFunc("/api/v1/tickets", authMiddleware.Authenticate(ticketHandler.ListTickets))
	mux.HandleFunc("/api/v1/tickets/create", authMiddleware.Authenticate(ticketHandler.CreateTicket))
	mux.HandleFunc("/api/v1/tickets/status", authMiddleware.Authenticate(ticketHandler.UpdateTicketStatus))

	// 5. Timesheet Logger Endpoints (Otentikasi Pengguna)
	mux.HandleFunc("/api/v1/timesheets", authMiddleware.Authenticate(timesheetHandler.ListTimesheets))
	mux.HandleFunc("/api/v1/timesheets/log", authMiddleware.Authenticate(timesheetHandler.LogWork))
	mux.HandleFunc("/api/v1/timesheets/summary", authMiddleware.Authenticate(timesheetHandler.GetSummary))

	// 6. Invoices & Billing Endpoints (Otentikasi Pengguna)
	mux.HandleFunc("/api/v1/invoices", authMiddleware.Authenticate(invoiceHandler.ListInvoices))
	mux.HandleFunc("/api/v1/invoices/create", adminOnly(invoiceHandler.CreateInvoice))
	mux.HandleFunc("/api/v1/invoices/status", adminOnly(invoiceHandler.UpdateInvoiceStatus))

	// 7. SAK EMKM Accounting & General Ledger Endpoints
	accountingStaff := authMiddleware.RequireRoles(domain.RoleDirektur, domain.RoleAdmin, domain.RoleAudit)
	mux.HandleFunc("/api/v1/journals", accountingStaff(accountingHandler.ListJournals))
	mux.HandleFunc("/api/v1/journals/create", adminOnly(accountingHandler.CreateJournal))

	// 8. Inbound Leads CRM Pipeline Endpoints
	if leadHandler != nil {
		mux.HandleFunc("/api/v1/leads/public", leadHandler.CreatePublicLead)
		mux.HandleFunc("/api/v1/leads", adminOnly(leadHandler.ListLeads))
		mux.HandleFunc("/api/v1/leads/status", adminOnly(leadHandler.UpdateLeadStatus))
	}

	// Wrap with Global Middlewares (Recovery, CORS, Logging)
	return withRecovery(withCORS(withLogging(mux)))
}

func withRecovery(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		defer func() {
			if rec := recover(); rec != nil {
				log.Printf("🔥 Internal Server Panic: %v", rec)
				w.Header().Set("Content-Type", "application/json")
				w.WriteHeader(http.StatusInternalServerError)
				_ = json.NewEncoder(w).Encode(map[string]interface{}{
					"success": false,
					"error":   fmt.Sprintf("Internal Server Error: %v", rec),
				})
			}
		}()
		next.ServeHTTP(w, r)
	})
}

func withCORS(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		origin := r.Header.Get("Origin")
		// Izinkan subdomain meldir.id dan localhost saat pengujian
		if origin != "" && (strings.HasSuffix(origin, "meldir.id") || strings.Contains(origin, "localhost")) {
			w.Header().Set("Access-Control-Allow-Origin", origin)
		} else {
			w.Header().Set("Access-Control-Allow-Origin", "*")
		}

		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS, PATCH")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization, X-Requested-With, Accept")
		w.Header().Set("Access-Control-Allow-Credentials", "true")

		if r.Method == http.MethodOptions {
			w.WriteHeader(http.StatusNoContent)
			return
		}

		next.ServeHTTP(w, r)
	})
}

func withLogging(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		next.ServeHTTP(w, r)
	})
}
