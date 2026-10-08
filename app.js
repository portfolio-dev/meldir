/**
 * Main Logic for meldir.id
 * Handles Bilingual translation, Portfolio filtering, Mobile navigation, PWA, and Active scroll indicators.
 */

document.addEventListener('DOMContentLoaded', () => {
    
    // --- 1. Language Toggle System ---
    const langIdBtn = document.getElementById('lang-id-btn');
    const langEnBtn = document.getElementById('lang-en-btn');
    const htmlTag = document.documentElement;

    // Translation Data for WhatsApp templates and attributes
    const waTemplates = {
        hero: {
            id: 'https://wa.me/628213173357?text=Halo%20meldir.id,%20saya%20tertarik%20untuk%20berdiskusi%20mengenai%20pembuatan%20aplikasi/website.',
            en: 'https://wa.me/628213173357?text=Hello%20meldir.id,%20I%20am%20interested%20in%20discussing%20a%20project%20for%20a%20website/app.'
        },
        social: {
            id: 'https://wa.me/628213173357?text=Halo%20meldir.id,%20saya%20dari%20yayasan%20ingin%20mengajukan%20kerjasama%20layanan%20pengembangan%20aplikasi%20sosial.',
            en: 'https://wa.me/628213173357?text=Hello%20meldir.id,%20I%20am%20from%20a%20foundation%20and%20would%20like%20to%20apply%20for%20the%20social%20app%20development%20program.'
        },
        quick: {
            id: 'https://wa.me/628213173357?text=Halo%20meldir.id,%20saya%20ingin%20konsultasi%20melalui%20Akses%20Cepat.',
            en: 'https://wa.me/628213173357?text=Hello%20meldir.id,%20I%20would%20like%20to%20consult%20via%20Quick%20Access.'
        }
    };

    // Safe localStorage wrapper to prevent crashes on file:/// protocol
    function getSavedLanguage() {
        try {
            return localStorage.getItem('preferred-language') || 'id';
        } catch (e) {
            return 'id'; // Default fallback
        }
    }

    function savePreferredLanguage(lang) {
        try {
            localStorage.setItem('preferred-language', lang);
        } catch (e) {
            // Silently ignore if localStorage is disabled
        }
    }

    function setLanguage(lang) {
        // Save to local storage
        savePreferredLanguage(lang);
        htmlTag.setAttribute('lang', lang);

        // Update active class on buttons
        if (lang === 'id') {
            langIdBtn.classList.add('active');
            langEnBtn.classList.remove('active');
        } else {
            langIdBtn.classList.remove('active');
            langEnBtn.classList.add('active');
        }

        // Translate elements with data-id and data-en
        const translatableElements = document.querySelectorAll('[data-id][data-en]');
        translatableElements.forEach(el => {
            const translation = el.getAttribute(`data-${lang}`);
            if (translation) {
                // If it's a list/nesting element or has simple tags, innerHTML is fine.
                // For simplicity and safety, if it does not contain HTML tags, use textContent.
                if (translation.includes('<') && translation.includes('>')) {
                    el.innerHTML = translation;
                } else {
                    el.textContent = translation;
                }
            }
        });

        // Translate document.title if present
        const titleEl = document.querySelector('title[data-id][data-en]');
        if (titleEl) {
            const titleTranslation = titleEl.getAttribute(`data-${lang}`);
            if (titleTranslation) {
                document.title = titleTranslation;
            }
        }

        // Translate dynamic URLs with data-href-id and data-href-en
        const translatableLinks = document.querySelectorAll('[data-href-id][data-href-en]');
        translatableLinks.forEach(link => {
            const href = link.getAttribute(`data-href-${lang}`);
            if (href) {
                link.setAttribute('href', href);
            }
        });
    }

    // Initialize Language
    const savedLang = getSavedLanguage();
    setLanguage(savedLang);

    // Event Listeners for Language Buttons
    langIdBtn.addEventListener('click', () => setLanguage('id'));
    langEnBtn.addEventListener('click', () => setLanguage('en'));


    // --- 2. Mobile Navigation System ---
    const mobileToggle = document.getElementById('mobile-toggle');
    const mobileNavOverlay = document.getElementById('mobile-nav-overlay');
    const mobileNavLinks = document.querySelectorAll('.mobile-nav-item');

    function toggleMobileMenu() {
        if (!mobileToggle || !mobileNavOverlay) return;
        const isOpen = mobileToggle.classList.toggle('open');
        mobileNavOverlay.classList.toggle('open', isOpen);
        document.body.style.overflow = isOpen ? 'hidden' : '';
    }

    function closeMobileMenu() {
        if (!mobileToggle || !mobileNavOverlay) return;
        mobileToggle.classList.remove('open');
        mobileNavOverlay.classList.remove('open');
        document.body.style.overflow = '';
    }

    if (mobileToggle) {
        mobileToggle.addEventListener('click', toggleMobileMenu);
    }
    
    mobileNavLinks.forEach(link => {
        link.addEventListener('click', closeMobileMenu);
    });

    // Close mobile menu on resize if screen gets larger
    window.addEventListener('resize', () => {
        if (window.innerWidth > 900) {
            closeMobileMenu();
        }
    });


    // --- 3. Scroll Header Styles & Active Menu ---
    const header = document.getElementById('site-header');
    const navItems = document.querySelectorAll('.nav-item');
    const sections = document.querySelectorAll('section[id]');

    function handleScroll() {
        if (!header) return;
        const scrollY = window.scrollY;

        // Scrolled header background style toggle
        if (scrollY > 20) {
            header.classList.add('scrolled');
        } else {
            header.classList.remove('scrolled');
        }

        // Active Navigation Highlight on Scroll
        let currentSectionId = '';
        sections.forEach(section => {
            const sectionHeight = section.offsetHeight;
            const sectionTop = section.offsetTop - 120; // offset header
            
            if (scrollY >= sectionTop && scrollY < sectionTop + sectionHeight) {
                currentSectionId = section.getAttribute('id');
            }
        });

        navItems.forEach(item => {
            item.classList.remove('active');
            if (item.getAttribute('href') === `#${currentSectionId}`) {
                item.classList.add('active');
            }
        });
    }

    window.addEventListener('scroll', handleScroll);
    handleScroll(); // Trigger immediately to set correct states on page load


    // --- 4. Security Modal Triggers ---
    const securityCards = document.querySelectorAll('.security-pillar-card');
    const securityModal = document.getElementById('security-modal');
    const closeSecurityModalBtn = document.getElementById('close-security-modal-btn');
    const heroSecurityBtn = document.querySelector('.hero-cta-group a[href="#security"]');

    if (securityModal && closeSecurityModalBtn) {
        const openSecurityModal = (pillarId) => {
            // Hide all detail contents
            const contents = document.querySelectorAll('.security-detail-content');
            contents.forEach(content => content.classList.add('hidden'));

            // Show active detail content
            const activeContent = document.getElementById(`security-content-${pillarId}`);
            if (activeContent) {
                activeContent.classList.remove('hidden');
            }

            // Open Modal
            securityModal.classList.remove('hidden');
            setTimeout(() => {
                securityModal.classList.add('open');
                document.body.style.overflow = 'hidden';
            }, 10);
        };

        const closeSecurityModal = () => {
            securityModal.classList.remove('open');
            setTimeout(() => {
                securityModal.classList.add('hidden');
                document.body.style.overflow = '';
            }, 300);
        };

        // Add event listeners to cards
        securityCards.forEach(card => {
            card.addEventListener('click', () => {
                const pillarId = card.getAttribute('data-pillar');
                openSecurityModal(pillarId);
            });
        });

        closeSecurityModalBtn.addEventListener('click', closeSecurityModal);

        if (heroSecurityBtn) {
            heroSecurityBtn.addEventListener('click', (e) => {
                e.preventDefault();
                const targetSection = document.getElementById('security');
                if (targetSection) {
                    targetSection.scrollIntoView({ behavior: 'smooth' });
                }
            });
        }

        securityModal.addEventListener('click', (e) => {
            if (e.target === securityModal) {
                closeSecurityModal();
            }
        });

        document.addEventListener('keydown', (e) => {
            if (e.key === 'Escape' && !securityModal.classList.contains('hidden')) {
                closeSecurityModal();
            }
        });
    }

    // --- 4b. Features Filter Logic ---
    const featureFilterBtns = document.querySelectorAll('.feature-filters .filter-btn');
    const featureCards = document.querySelectorAll('.feature-card');

    featureFilterBtns.forEach(btn => {
        btn.addEventListener('click', (e) => {
            // Remove active from all feature filter buttons and add to clicked
            featureFilterBtns.forEach(b => b.classList.remove('active'));
            e.target.classList.add('active');

            const filterValue = e.target.getAttribute('data-feature-filter');

            featureCards.forEach(card => {
                const cardCategory = card.getAttribute('data-feature-category');

                if (filterValue === 'all' || cardCategory === filterValue) {
                    card.classList.remove('hidden');
                    card.style.display = 'flex';
                    setTimeout(() => {
                        card.style.opacity = '1';
                        card.style.transform = 'scale(1)';
                    }, 50);
                } else {
                    card.style.opacity = '0';
                    card.style.transform = 'scale(0.92)';
                    setTimeout(() => {
                        if (card.style.opacity === '0') {
                            card.classList.add('hidden');
                            card.style.display = 'none';
                        }
                    }, 300);
                }
            });
        });
    });



    // --- 5. PWA Installation Setup ---
    let deferredPrompt;
    const pwaInstallBtn = document.getElementById('pwa-install-btn');

    if (pwaInstallBtn) {
        window.addEventListener('beforeinstallprompt', (e) => {
            // Prevent Chrome 67 and earlier from automatically showing the prompt
            e.preventDefault();
            // Stash the event so it can be triggered later.
            deferredPrompt = e;
            // Update UI to notify user they can install the PWA
            pwaInstallBtn.classList.remove('hidden');
        });

        pwaInstallBtn.addEventListener('click', async () => {
            if (!deferredPrompt) return;
            
            // Show the install prompt
            deferredPrompt.prompt();
            
            // Wait for the user to respond to the prompt
            const { outcome } = await deferredPrompt.userChoice;
            console.log(`User response to the install prompt: ${outcome}`);
            
            // We've used the prompt, and can't use it again, clear it
            deferredPrompt = null;
            // Hide the install button
            pwaInstallBtn.classList.add('hidden');
        });

        window.addEventListener('appinstalled', (evt) => {
            console.log('meldir.id app was installed.');
            pwaInstallBtn.classList.add('hidden');
        });
    }

    // --- 4c. Features Modal Open/Close ---
    const openFeaturesModalBtn = document.getElementById('open-features-modal-btn');
    const closeFeaturesModalBtn = document.getElementById('close-features-modal-btn');
    const featuresModal = document.getElementById('features-modal');

    if (openFeaturesModalBtn && closeFeaturesModalBtn && featuresModal) {
        const openModal = () => {
            featuresModal.classList.remove('hidden');
            setTimeout(() => {
                featuresModal.classList.add('open');
                document.body.style.overflow = 'hidden';
            }, 10);
        };

        const closeModal = () => {
            featuresModal.classList.remove('open');
            setTimeout(() => {
                featuresModal.classList.add('hidden');
                document.body.style.overflow = '';
            }, 300);
        };

        openFeaturesModalBtn.addEventListener('click', openModal);
        closeFeaturesModalBtn.addEventListener('click', closeModal);

        featuresModal.addEventListener('click', (e) => {
            if (e.target === featuresModal) {
                closeModal();
            }
        });

        document.addEventListener('keydown', (e) => {
            if (e.key === 'Escape' && !featuresModal.classList.contains('hidden')) {
                closeModal();
            }
        });
    }

    // --- 4d. Security Modal Cleanup (Handled in 4.) ---

    // --- 4e. Devices Mockup Slider Dots Synchronization (Mobile) ---
    const showcaseContainer = document.querySelector('.devices-showcase-container');
    const sliderDots = document.querySelectorAll('.device-slider-dots .slider-dot');

    if (showcaseContainer && sliderDots.length > 0) {
        showcaseContainer.addEventListener('scroll', () => {
            const index = Math.round(showcaseContainer.scrollLeft / showcaseContainer.clientWidth);
            sliderDots.forEach((dot, i) => {
                dot.classList.toggle('active', i === index);
            });
        });

        sliderDots.forEach((dot, i) => {
            dot.addEventListener('click', () => {
                showcaseContainer.scrollTo({
                    left: i * showcaseContainer.clientWidth,
                    behavior: 'smooth'
                });
            });
        });
    }


    // --- 6. BCA-Style Hero Carousel Slider ---
    const heroTrack = document.getElementById('hero-slider-track');
    const heroSlides = document.querySelectorAll('.hero-slide');
    const heroPrevBtn = document.getElementById('hero-prev-btn');
    const heroNextBtn = document.getElementById('hero-next-btn');
    const heroDots = document.querySelectorAll('.hero-slider-dots .hero-dot');
    const heroSliderSection = document.getElementById('hero');

    if (heroTrack && heroSlides.length > 0) {
        let currentSlideIndex = 0;
        const totalSlides = heroSlides.length;
        let slideInterval = null;

        function updateSlidePosition() {
            heroTrack.style.transform = `translateX(-${currentSlideIndex * 100}%)`;
            
            heroSlides.forEach((slide, idx) => {
                slide.classList.toggle('active', idx === currentSlideIndex);
            });

            heroDots.forEach((dot, idx) => {
                dot.classList.toggle('active', idx === currentSlideIndex);
            });
        }

        function nextSlide() {
            currentSlideIndex = (currentSlideIndex + 1) % totalSlides;
            updateSlidePosition();
        }

        function prevSlide() {
            currentSlideIndex = (currentSlideIndex - 1 + totalSlides) % totalSlides;
            updateSlidePosition();
        }

        function goToSlide(index) {
            currentSlideIndex = index;
            updateSlidePosition();
        }

        // Event listeners for prev/next buttons
        if (heroPrevBtn) {
            heroPrevBtn.addEventListener('click', () => {
                prevSlide();
                restartAutoSlide();
            });
        }

        if (heroNextBtn) {
            heroNextBtn.addEventListener('click', () => {
                nextSlide();
                restartAutoSlide();
            });
        }

        // Event listeners for dots
        heroDots.forEach((dot, idx) => {
            dot.addEventListener('click', () => {
                goToSlide(idx);
                restartAutoSlide();
            });
        });

        // Auto-play every 6 seconds
        function startAutoSlide() {
            if (!slideInterval) {
                slideInterval = setInterval(nextSlide, 6000);
            }
        }

        function stopAutoSlide() {
            if (slideInterval) {
                clearInterval(slideInterval);
                slideInterval = null;
            }
        }

        function restartAutoSlide() {
            stopAutoSlide();
            startAutoSlide();
        }

        startAutoSlide();

        // Pause on mouse enter, resume on mouse leave
        if (heroSliderSection) {
            heroSliderSection.addEventListener('mouseenter', stopAutoSlide);
            heroSliderSection.addEventListener('mouseleave', startAutoSlide);

            // Touch Swipe Detection for mobile devices
            let touchStartX = 0;
            let touchEndX = 0;

            heroSliderSection.addEventListener('touchstart', (e) => {
                touchStartX = e.changedTouches[0].screenX;
                stopAutoSlide();
            }, { passive: true });

            heroSliderSection.addEventListener('touchend', (e) => {
                touchEndX = e.changedTouches[0].screenX;
                const swipeThreshold = 40;
                if (touchStartX - touchEndX > swipeThreshold) {
                    nextSlide();
                } else if (touchEndX - touchStartX > swipeThreshold) {
                    prevSlide();
                }
                startAutoSlide();
            }, { passive: true });
        }
    }


    // --- 7. Full-Height BCA Sticky Sidebar Scroll System (Desktop Only) ---
    const quickSidebar = document.getElementById('bca-quick-sidebar');
    const heroSection = document.getElementById('hero');

    if (quickSidebar) {
        function updateSidebarByScroll() {
            // Disabled in cockpit mode (uses floating dock with hover flyouts)
            if (document.body.classList.contains('cockpit-locked')) {
                document.body.classList.remove('sidebar-expanded');
                quickSidebar.classList.remove('expanded');
                return;
            }

            // Hanya aktif untuk tampilan desktop (> 1024px)
            if (window.innerWidth <= 1024) {
                document.body.classList.remove('sidebar-expanded');
                quickSidebar.classList.remove('expanded');
                return;
            }

            const scrollY = window.scrollY;
            const heroHeight = heroSection ? heroSection.offsetHeight : 600;

            // Selama layar menampilkan kondisi di section atas/hero: TERBUKA PENUH (sidebar-expanded)
            // Tertutup otomatis meninggalkan icon saja jika scroll ke bawah melewati hero
            // Kembali terbuka penuh jika naik ke section hero
            if (scrollY < heroHeight - 120) {
                document.body.classList.add('sidebar-expanded');
                quickSidebar.classList.add('expanded');
            } else {
                document.body.classList.remove('sidebar-expanded');
                quickSidebar.classList.remove('expanded');
            }
        }

        window.addEventListener('scroll', updateSidebarByScroll, { passive: true });
        window.addEventListener('resize', updateSidebarByScroll, { passive: true });
        updateSidebarByScroll(); // Inisialisasi saat pertama dimuat
    }

    // --- 8. Cockpit No-Scroll & Universal Drawer Controller ---
    const cockpitBackdrop = document.getElementById('cockpit-drawer-backdrop');
    const cockpitDrawers = document.querySelectorAll('.cockpit-drawer');
    const drawerTriggers = document.querySelectorAll('[data-drawer]');
    const drawerCloseBtns = document.querySelectorAll('.drawer-close-btn, [data-close-drawer]');
    const segmentBtns = document.querySelectorAll('.cockpit-segment-btn');
    const overviewBtns = document.querySelectorAll('[data-target="overview"]');

    function openCockpitDrawer(drawerId) {
        if (!drawerId) return;
        const targetDrawer = document.getElementById(drawerId);
        if (!targetDrawer) return;

        // If target drawer is already active, toggle it closed
        if (targetDrawer.classList.contains('active')) {
            closeCockpitDrawer();
            return;
        }

        // Close any other open drawer
        cockpitDrawers.forEach(d => d.classList.remove('active'));

        // Open target drawer & backdrop
        targetDrawer.classList.add('active');
        if (cockpitBackdrop) cockpitBackdrop.classList.add('active');

        // Update active states on segment buttons & mobile tabs
        segmentBtns.forEach(btn => {
            btn.classList.toggle('active', btn.getAttribute('data-drawer') === drawerId);
        });
        mobileTabItems.forEach(tab => {
            tab.classList.toggle('active', tab.getAttribute('data-drawer') === drawerId);
        });
    }

    function closeCockpitDrawer() {
        cockpitDrawers.forEach(d => d.classList.remove('active'));
        if (cockpitBackdrop) cockpitBackdrop.classList.remove('active');

        // Reset segment button active state to overview
        segmentBtns.forEach(btn => {
            btn.classList.toggle('active', btn.getAttribute('data-target') === 'overview');
        });
        mobileTabItems.forEach(tab => {
            tab.classList.toggle('active', tab.getAttribute('data-target') === 'overview');
        });

        // Clear hash without reload
        if (window.location.hash) {
            history.pushState(null, document.title, window.location.pathname + window.location.search);
        }
    }

    // Event listeners for triggers
    drawerTriggers.forEach(trigger => {
        trigger.addEventListener('click', (e) => {
            e.preventDefault();
            const drawerId = trigger.getAttribute('data-drawer');
            openCockpitDrawer(drawerId);
        });
    });

    // Overview buttons (Ikhtisar / Beranda)
    overviewBtns.forEach(btn => {
        btn.addEventListener('click', (e) => {
            e.preventDefault();
            closeCockpitDrawer();
        });
    });

    // Close buttons & Backdrop
    drawerCloseBtns.forEach(btn => btn.addEventListener('click', closeCockpitDrawer));
    if (cockpitBackdrop) cockpitBackdrop.addEventListener('click', closeCockpitDrawer);

    // Keyboard ESC to close drawer
    document.addEventListener('keydown', (e) => {
        if (e.key === 'Escape') closeCockpitDrawer();
    });

    // Hash routing on load (e.g. #services, #security, #workflow, #comparison, #social)
    function checkHashRoute() {
        const hash = window.location.hash.toLowerCase().replace('#', '');
        if (!hash) return;
        const hashMap = {
            'services': 'drawer-services',
            'service': 'drawer-services',
            'security': 'drawer-security',
            'privacy': 'drawer-security',
            'legal': 'drawer-security',
            'workflow': 'drawer-workflow',
            'cara-kerja': 'drawer-workflow',
            'comparison': 'drawer-comparison',
            'social': 'drawer-social',
            'audit': 'drawer-audit'
        };
        if (hashMap[hash]) {
            openCockpitDrawer(hashMap[hash]);
        }
    }
    // --- 9. Mobile Metrics Row Swipe & Drag Controller ---
    const metricsRow = document.querySelector('.cockpit-metrics-row');
    if (metricsRow) {
        let isDown = false;
        let startX = 0;
        let scrollLeft = 0;

        // Mouse Drag Support (Desktop & Emulation)
        metricsRow.addEventListener('mousedown', (e) => {
            isDown = true;
            metricsRow.classList.add('dragging');
            startX = e.pageX - metricsRow.offsetLeft;
            scrollLeft = metricsRow.scrollLeft;
        });

        metricsRow.addEventListener('mouseleave', () => {
            isDown = false;
            metricsRow.classList.remove('dragging');
        });

        metricsRow.addEventListener('mouseup', () => {
            isDown = false;
            metricsRow.classList.remove('dragging');
        });

        metricsRow.addEventListener('mousemove', (e) => {
            if (!isDown) return;
            e.preventDefault();
            const x = e.pageX - metricsRow.offsetLeft;
            const walk = (x - startX) * 1.5;
            metricsRow.scrollLeft = scrollLeft - walk;
        });

        // Touch Swipe & Drag Support (Mobile & Touch Devices)
        let touchStartX = 0;
        let touchStartScroll = 0;
        metricsRow.addEventListener('touchstart', (e) => {
            if (e.touches && e.touches.length === 1) {
                touchStartX = e.touches[0].pageX;
                touchStartScroll = metricsRow.scrollLeft;
            }
        }, { passive: true });

        metricsRow.addEventListener('touchmove', (e) => {
            if (!touchStartX || !e.touches || e.touches.length !== 1) return;
            const currentX = e.touches[0].pageX;
            const diff = currentX - touchStartX;
            metricsRow.scrollLeft = touchStartScroll - diff;
        }, { passive: true });

        metricsRow.addEventListener('touchend', () => {
            touchStartX = 0;
        }, { passive: true });
    }

    // --- 10. Lead Audit Modal & Inbound CRM Form Handler ---
    const leadModal = document.getElementById('lead-audit-modal');
    const leadModalClose = document.getElementById('lead-modal-close-btn');
    const leadTriggers = document.querySelectorAll('[data-modal="lead-audit-modal"]');
    const leadForm = document.getElementById('lead-audit-form');
    const leadSuccess = document.getElementById('lead-success-container');
    const leadCodeDisplay = document.getElementById('lead-code-display');
    const leadWaLink = document.getElementById('lead-wa-direct-link');
    const leadError = document.getElementById('lead-form-error');
    const leadSubmitBtn = document.getElementById('lead-submit-btn');

    function openLeadModal() {
        if (!leadModal) return;
        leadModal.classList.add('active');
        document.body.style.overflow = 'hidden';
    }

    function closeLeadModal() {
        if (!leadModal) return;
        leadModal.classList.remove('active');
        document.body.style.overflow = '';
    }

    leadTriggers.forEach(btn => {
        btn.addEventListener('click', (e) => {
            e.preventDefault();
            openLeadModal();
        });
    });

    if (leadModalClose) {
        leadModalClose.addEventListener('click', closeLeadModal);
    }

    if (leadModal) {
        leadModal.addEventListener('click', (e) => {
            if (e.target === leadModal) {
                closeLeadModal();
            }
        });
    }

    document.addEventListener('keydown', (e) => {
        if (e.key === 'Escape' && leadModal && leadModal.classList.contains('active')) {
            closeLeadModal();
        }
    });

    if (leadForm) {
        leadForm.addEventListener('submit', async (e) => {
            e.preventDefault();
            if (leadError) {
                leadError.style.display = 'none';
                leadError.textContent = '';
            }

            const name = document.getElementById('lead-name')?.value.trim() || '';
            const company = document.getElementById('lead-company')?.value.trim() || '';
            const whatsapp = document.getElementById('lead-whatsapp')?.value.trim() || '';
            const email = document.getElementById('lead-email')?.value.trim() || '';
            const service = document.getElementById('lead-service')?.value || 'Sistem Baru (Web/Mobile)';
            const budget = document.getElementById('lead-budget')?.value || 'Fleksibel';
            const notes = document.getElementById('lead-notes')?.value.trim() || '';

            if (!name || !whatsapp) {
                if (leadError) {
                    leadError.textContent = 'Mohon lengkapi Nama Lengkap dan Nomor WhatsApp Anda.';
                    leadError.style.display = 'block';
                }
                return;
            }

            if (leadSubmitBtn) {
                leadSubmitBtn.disabled = true;
                leadSubmitBtn.innerHTML = '<span>Mengirim permohonan...</span>';
            }

            const payload = {
                name,
                company,
                whatsapp,
                email,
                service_interest: service,
                budget_range: budget,
                notes,
                source: 'meldir.id-landing-audit'
            };

            let leadCode = 'LEAD-' + new Date().getFullYear() + '-' + Math.floor(100 + Math.random() * 900);

            try {
                let targetEndpoint = 'https://office.meldir.id/api/v1/leads/public';
                if (window.location.hostname === 'office.meldir.id' || window.location.hostname === 'localhost') {
                    targetEndpoint = '/api/v1/leads/public';
                }

                const response = await fetch(targetEndpoint, {
                    method: 'POST',
                    headers: { 'Content-Type': 'application/json' },
                    body: JSON.stringify(payload)
                });

                if (response.ok) {
                    const data = await response.json();
                    if (data && (data.lead_code || (data.data && data.data.lead_code))) {
                        leadCode = data.lead_code || data.data.lead_code;
                    }
                } else {
                    console.warn('API returned status:', response.status);
                }
            } catch (err) {
                console.warn('Backend connection notice:', err);
            }

            // Display success view
            if (leadCodeDisplay) {
                leadCodeDisplay.textContent = leadCode;
            }

            if (leadWaLink) {
                const waMessage = `Halo Meldir, saya mengajukan Audit Sistem & Konsultasi melalui website:\n\n` +
                    `*Kode Tiket:* ${leadCode}\n` +
                    `*Nama:* ${name}\n` +
                    `*Perusahaan:* ${company || '-'}\n` +
                    `*Layanan:* ${service}\n` +
                    `*Estimasi Budget:* ${budget}\n` +
                    (notes ? `*Catatan/Kendala:* ${notes}\n\n` : '\n') +
                    `Mohon arahan dan telaah teknisnya. Terima kasih!`;
                leadWaLink.href = `https://wa.me/628213173357?text=${encodeURIComponent(waMessage)}`;
            }

            leadForm.style.display = 'none';
            if (leadSuccess) {
                leadSuccess.style.display = 'block';
            }
        });
    }

    checkHashRoute();
    window.addEventListener('hashchange', checkHashRoute);

    if (window.location.hash.toLowerCase() === '#audit') {
        openLeadModal();
    }
    if (window.location.hash.toLowerCase() === '#chat') {
        openLiveChat();
    }

    // --- 11. Meldir Live Chat Online System with Pre-Chat Onboarding ---
    const chatWrapper = document.getElementById('meldir-chat-widget');
    const chatLauncher = document.getElementById('chat-launcher-btn');
    const chatMinimize = document.getElementById('chat-btn-minimize');
    const chatBtnReset = document.getElementById('chat-btn-reset');
    const chatOnboardingView = document.getElementById('chat-onboarding-view');
    const chatActiveView = document.getElementById('chat-active-view');
    const chatPreForm = document.getElementById('chat-pre-form');
    const chatMessages = document.getElementById('chat-messages-container');
    const chatForm = document.getElementById('chat-input-form');
    const chatInput = document.getElementById('chat-text-input');
    const chatConnectedName = document.getElementById('chat-connected-name');
    const chatSessionBadge = document.getElementById('chat-session-badge');
    const btnSubmitStartChat = document.getElementById('btn-submit-start-chat');
    const chatTriggers = document.querySelectorAll('[data-open-chat]');
    const chatChips = document.querySelectorAll('.chat-chip-btn');

    let chatPollingTimer = null;
    let activeSession = null;
    const renderedMsgIds = new Set();

    function loadSavedSession() {
        try {
            const raw = sessionStorage.getItem('meldir_chat_session');
            if (raw) {
                return JSON.parse(raw);
            }
        } catch (e) {
            console.error('Gagal membaca sesi chat:', e);
        }
        return null;
    }

    function saveSession(session) {
        activeSession = session;
        sessionStorage.setItem('meldir_chat_session', JSON.stringify(session));
    }

    function clearSession() {
        activeSession = null;
        renderedMsgIds.clear();
        sessionStorage.removeItem('meldir_chat_session');
        if (chatPollingTimer) {
            clearInterval(chatPollingTimer);
            chatPollingTimer = null;
        }
        showOnboardingView();
    }

    function showOnboardingView(preselectedTopic = null) {
        if (chatOnboardingView) chatOnboardingView.style.display = 'flex';
        if (chatActiveView) chatActiveView.style.display = 'none';
        if (chatBtnReset) chatBtnReset.style.display = 'none';

        if (preselectedTopic) {
            const topicSelect = document.getElementById('chat-input-topic');
            if (topicSelect) {
                for (let i = 0; i < topicSelect.options.length; i++) {
                    if (topicSelect.options[i].text.toLowerCase().includes(preselectedTopic.toLowerCase())) {
                        topicSelect.selectedIndex = i;
                        break;
                    }
                }
            }
        }
    }

    function showActiveView(session) {
        if (chatOnboardingView) chatOnboardingView.style.display = 'none';
        if (chatActiveView) chatActiveView.style.display = 'flex';
        if (chatBtnReset) chatBtnReset.style.display = 'inline-flex';

        if (chatConnectedName && session.visitor_name) {
            chatConnectedName.textContent = session.visitor_name;
        }
        if (chatSessionBadge && session.session_code) {
            chatSessionBadge.textContent = session.session_code;
        }

        // Poll messages
        fetchMessages(session.session_code);
        if (!chatPollingTimer) {
            chatPollingTimer = setInterval(() => {
                if (activeSession && chatWrapper?.classList.contains('active')) {
                    fetchMessages(activeSession.session_code);
                }
            }, 3500);
        }
    }

    function appendMessageUI(msg) {
        if (!chatMessages) return;
        if (msg.id && renderedMsgIds.has(msg.id)) return;
        if (msg.id) renderedMsgIds.add(msg.id);

        const isVisitor = msg.sender_type === 'visitor';
        const msgEl = document.createElement('div');
        msgEl.className = `chat-msg ${isVisitor ? 'msg-user' : 'msg-bot'}`;

        const bubbleEl = document.createElement('div');
        bubbleEl.className = 'msg-bubble';

        if (!isVisitor && msg.sender_name) {
            const senderTag = document.createElement('div');
            senderTag.style.fontSize = '0.7rem';
            senderTag.style.fontWeight = '700';
            senderTag.style.color = '#0060af';
            senderTag.style.marginBottom = '4px';
            senderTag.textContent = msg.sender_name;
            bubbleEl.appendChild(senderTag);
        }

        const p = document.createElement('p');
        p.textContent = msg.message;
        bubbleEl.appendChild(p);

        const timeEl = document.createElement('span');
        timeEl.className = 'msg-time';
        if (msg.created_at) {
            try {
                const d = new Date(msg.created_at);
                timeEl.textContent = `${d.getHours().toString().padStart(2, '0')}:${d.getMinutes().toString().padStart(2, '0')}`;
            } catch (e) {
                timeEl.textContent = 'Baru saja';
            }
        } else {
            timeEl.textContent = 'Baru saja';
        }

        msgEl.appendChild(bubbleEl);
        msgEl.appendChild(timeEl);
        chatMessages.appendChild(msgEl);
        chatMessages.scrollTop = chatMessages.scrollHeight;
    }

    async function fetchMessages(sessionCode) {
        if (!sessionCode) return;
        try {
            const res = await fetch(`/api/v1/chat/messages?session_code=${encodeURIComponent(sessionCode)}`);
            if (!res.ok) return;
            const data = await res.json();
            if (data.success && Array.isArray(data.messages)) {
                data.messages.forEach(m => appendMessageUI(m));
            }
        } catch (e) {
            console.warn('Gagal memuat pesan live chat:', e);
        }
    }

    // Initialize session state on page load
    activeSession = loadSavedSession();
    if (activeSession && activeSession.session_code) {
        showActiveView(activeSession);
    } else {
        showOnboardingView();
    }

    function openLiveChat(topic = null) {
        if (!chatWrapper) return;
        chatWrapper.classList.add('active');

        if (!activeSession) {
            showOnboardingView(topic);
            const nameInput = document.getElementById('chat-input-name');
            if (nameInput) setTimeout(() => nameInput.focus(), 250);
        } else {
            showActiveView(activeSession);
            if (chatInput) setTimeout(() => chatInput.focus(), 250);
        }
    }

    function closeLiveChat() {
        if (!chatWrapper) return;
        chatWrapper.classList.remove('active');
    }

    function toggleLiveChat() {
        if (!chatWrapper) return;
        if (chatWrapper.classList.contains('active')) {
            closeLiveChat();
        } else {
            openLiveChat();
        }
    }

    if (chatLauncher) {
        chatLauncher.addEventListener('click', toggleLiveChat);
    }
    if (chatMinimize) {
        chatMinimize.addEventListener('click', closeLiveChat);
    }
    if (chatBtnReset) {
        chatBtnReset.addEventListener('click', () => {
            if (confirm('Apakah Anda ingin mengakhiri sesi chat ini dan memulai konsultasi baru?')) {
                clearSession();
            }
        });
    }

    chatTriggers.forEach(btn => {
        btn.addEventListener('click', (e) => {
            e.preventDefault();
            const topic = btn.getAttribute('data-chat-topic') || null;
            openLiveChat(topic);
        });
    });

    chatChips.forEach(chip => {
        chip.addEventListener('click', () => {
            const query = chip.getAttribute('data-query') || chip.textContent.trim();
            if (activeSession && chatInput) {
                chatInput.value = query;
                chatInput.focus();
            }
        });
    });

    // Handle Pre-Chat Form Submit
    if (chatPreForm) {
        chatPreForm.addEventListener('submit', async (e) => {
            e.preventDefault();
            const name = (document.getElementById('chat-input-name')?.value || '').trim();
            const phone = (document.getElementById('chat-input-phone')?.value || '').trim();
            const email = (document.getElementById('chat-input-email')?.value || '').trim();
            const topic = (document.getElementById('chat-input-topic')?.value || '').trim();
            const message = (document.getElementById('chat-input-initial-msg')?.value || '').trim();

            if (!name || !phone || !message) {
                alert('Silakan lengkapi Nama, No. WhatsApp/HP, dan Pertanyaan Awal Anda.');
                return;
            }

            if (btnSubmitStartChat) {
                btnSubmitStartChat.disabled = true;
                btnSubmitStartChat.innerHTML = '<span>⏳ Memulai Sesi Konsultasi...</span>';
            }

            try {
                const res = await fetch('/api/v1/chat/start', {
                    method: 'POST',
                    headers: { 'Content-Type': 'application/json' },
                    body: JSON.stringify({
                        name,
                        phone,
                        email,
                        service_interest: topic,
                        message
                    })
                });

                const data = await res.json();
                if (data.success && data.data) {
                    saveSession(data.data);
                    // Clear messages container
                    if (chatMessages) {
                        chatMessages.innerHTML = '<div class="chat-time-chip">Sesi Konsultasi Dimulai</div>';
                    }
                    renderedMsgIds.clear();
                    showActiveView(data.data);
                    if (Array.isArray(data.data.messages)) {
                        data.data.messages.forEach(m => appendMessageUI(m));
                    }
                } else {
                    alert('Gagal memulai sesi chat: ' + (data.error || 'Terjadi kesalahan sistem'));
                }
            } catch (err) {
                console.error('Error starting chat:', err);
                // Fallback offline mock session
                const mockSession = {
                    id: Date.now(),
                    session_code: 'CHAT-OFFLINE-' + Math.floor(Math.random() * 1000),
                    visitor_name: name,
                    visitor_phone: phone,
                    service_interest: topic
                };
                saveSession(mockSession);
                showActiveView(mockSession);
                appendMessageUI({
                    sender_type: 'visitor',
                    sender_name: name,
                    message: message,
                    created_at: new Date()
                });
                appendMessageUI({
                    sender_type: 'agent',
                    sender_name: 'Meldir Virtual Desk',
                    message: `Halo Bpk/Ibu ${name}! Pesan dan data Anda telah kami terima. Konsultan kami akan merespons sesegera mungkin di sini.`,
                    created_at: new Date()
                });
            } finally {
                if (btnSubmitStartChat) {
                    btnSubmitStartChat.disabled = false;
                    btnSubmitStartChat.innerHTML = '<span>🚀 Mulai Sesi Chat Online</span>';
                }
            }
        });
    }

    // Handle Active Conversation Message Submit
    if (chatForm) {
        chatForm.addEventListener('submit', async (e) => {
            e.preventDefault();
            const text = (chatInput?.value || '').trim();
            if (!text || !activeSession) return;

            // Clear input
            if (chatInput) chatInput.value = '';

            // Optimistic UI Append
            const tempMsg = {
                sender_type: 'visitor',
                sender_name: activeSession.visitor_name || 'Pengunjung',
                message: text,
                created_at: new Date()
            };
            appendMessageUI(tempMsg);

            try {
                const res = await fetch('/api/v1/chat/message', {
                    method: 'POST',
                    headers: { 'Content-Type': 'application/json' },
                    body: JSON.stringify({
                        session_code: activeSession.session_code,
                        sender_name: activeSession.visitor_name,
                        message: text
                    })
                });
                const data = await res.json();
                if (data.success && data.data && data.data.id) {
                    renderedMsgIds.add(data.data.id);
                }
            } catch (err) {
                console.warn('Gagal mengirim pesan chat:', err);
            }
        });
    }


});

// --- Register Service Worker ---
if ('serviceWorker' in navigator) {
    window.addEventListener('load', () => {
        navigator.serviceWorker.register('sw.js')
            .then(reg => console.log('Service Worker registered successfully.', reg.scope))
            .catch(err => console.log('Service Worker registration failed: ', err));
    });
}

