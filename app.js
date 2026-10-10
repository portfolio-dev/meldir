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
    const mobileTabItems = document.querySelectorAll('.cockpit-tab-item');
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

    const leadCloseBtns = document.querySelectorAll('#lead-modal-close-btn, [data-close-lead-modal], #lead-audit-modal .lead-modal-close');
    leadCloseBtns.forEach(btn => {
        btn.addEventListener('click', closeLeadModal);
    });

    if (leadModal) {
        leadModal.addEventListener('click', (e) => {
            if (e.target === leadModal) {
                closeLeadModal();
            }
        });
    }

    document.addEventListener('keydown', (e) => {
        if (e.key === 'Escape') {
            if (leadModal && leadModal.classList.contains('active')) closeLeadModal();
            if (slaModal && slaModal.classList.contains('active')) closeSlaModal();
            if (secModal && secModal.classList.contains('active')) closeSecurityDetailModal();
            if (typeof closeAllGadgetModals === 'function') closeAllGadgetModals();
        }
    });

    // --- 10B. SLA 99.9% Educational Modal Controller ---
    const slaModal = document.getElementById('sla-detail-modal');
    const slaTriggers = document.querySelectorAll('[data-modal="sla-detail-modal"]');
    const slaCloseBtns = document.querySelectorAll('[data-close-sla-modal]');

    function openSlaModal() {
        if (!slaModal) return;
        slaModal.classList.add('active');
        document.body.style.overflow = 'hidden';
    }

    function closeSlaModal() {
        if (!slaModal) return;
        slaModal.classList.remove('active');
        document.body.style.overflow = '';
    }

    slaTriggers.forEach(btn => {
        btn.addEventListener('click', (e) => {
            e.preventDefault();
            openSlaModal();
        });
    });

    slaCloseBtns.forEach(btn => {
        btn.addEventListener('click', closeSlaModal);
    });

    if (slaModal) {
        slaModal.addEventListener('click', (e) => {
            if (e.target === slaModal) closeSlaModal();
        });
    }

    // --- 10C. 12 Standar Keamanan & Legalitas Educational Modal Controller ---
    const secModal = document.getElementById('security-detail-modal');
    const secCloseBtns = document.querySelectorAll('[data-close-sec-modal]');
    const secModalCat = document.getElementById('sec-modal-category');
    const secModalBadge = document.getElementById('sec-modal-badge');
    const secModalTitle = document.getElementById('sec-modal-title');
    const secModalDesc = document.getElementById('sec-modal-desc');
    const secModalImpl = document.getElementById('sec-modal-impl');
    const secModalBenefit = document.getElementById('sec-modal-benefit');
    const secModalTip = document.getElementById('sec-modal-tip');
    const secModalChatBtn = document.getElementById('sec-modal-chat-btn');

    const securityStandardsData = {
        'ssl-grade-a': {
            category: 'Keamanan Data & Jaringan',
            title: 'Enkripsi SSL/TLS Grade A+',
            badge: 'Transmisi Data 256-bit',
            desc: 'Semua lalu lintas data antara browser pengunjung dan server aplikasi diacak menggunakan protokol enkripsi kriptografi TLS 1.3 standar perbankan. Ini memastikan tidak ada pihak ketiga di jaringan Wi-Fi publik, ISP, atau penyadap yang bisa mengintip (eavesdropping) atau menyadap data sensitif.',
            impl: 'Meldir mengonfigurasi cipher suite modern (ECDHE-ECDSA-AES256-GCM), mengaktifkan HSTS (HTTP Strict Transport Security), OCSP Stapling, dan Perfect Forward Secrecy (PFS), menghasilkan skor evaluasi keamanan Grade A+ resmi pada uji Qualys SSL Labs.',
            benefit: 'Mencegah pencurian kredensial akun, data kartu/pembayaran, dan manipulasi konten (Man-in-the-Middle attack). Membangun rasa aman seketika bagi pelanggan dengan ikon gembok aman di browser.',
            tip: 'Sertifikat SSL gratisan tanpa konfigurasi cipher yang tepat sering kali masih rentan terhadap downgrade attack ke protokol lama. Meldir memastikan proteksi TLS 1.3 murni tanpa celah warisan.',
            chatTopic: 'Konsultasi Enkripsi SSL Grade A+ & Keamanan Jaringan'
        },
        'hashing-password': {
            category: 'Kriptografi & Autentikasi',
            title: 'Hashing Kriptografis Password',
            badge: 'Bcrypt & Argon2 Kuat',
            desc: 'Kata sandi pengguna maupun staf internal TIDAK PERNAH disimpan dalam bentuk teks biasa (plain-text). Melalui proses hashing satu arah dengan garam acak (cryptographic salt), kata sandi diubah menjadi deretan karakter acak yang mustahil dikembalikan ke teks aslinya.',
            impl: 'Meldir menerapkan algoritma pemenang kompetisi kriptografi dunia yaitu Argon2id atau Bcrypt dengan work-factor (cost) terkalibrasi tinggi, mencegah peretasan via rainbow tables atau serangan brute-force berbasis GPU superkomputer.',
            benefit: 'Sekalipun terjadi skenario terburuk di mana database aplikasi bocor atau diretas, penyerang tetap tidak dapat membaca kata sandi pengguna atau pimpinan perusahaan Anda.',
            tip: 'Jangan pernah mempercayakan aplikasi bisnis pada vendor yang masih menggunakan MD5 atau SHA1, karena algoritma tersebut sudah dapat ditembus dalam hitungan detik menggunakan kamus peretas modern.',
            chatTopic: 'Konsultasi Standar Hashing Password & Autentikasi'
        },
        'cloud-backup': {
            category: 'Ketahanan Infrastruktur',
            title: 'Pencadangan Cloud Otomatis',
            badge: 'Snapshot Harian Terisolasi',
            desc: 'Sistem pencadangan otomatis tanpa henti yang menyalin seluruh basis data, aset media, dan konfigurasi server ke media penyimpanan awan (cloud storage) terpisah secara terjadwal setiap hari.',
            impl: 'Snapshot terenkripsi AES-256 dibuat otomatis setiap tengah malam dan disimpan pada bucket off-site terpisah dengan retention policy 30 hari. Uji integritas restorasi data dijalankan secara berkala untuk memastikan backup selalu valid.',
            benefit: 'Kekebalan mutlak dari risiko kehilangan data akibat kerusakan hardware server, kesalahan manusia (human error/salah hapus staf), serangan ransomware, atau kegagalan penyedia hosting.',
            tip: 'Banyak sistem memiliki backup tetapi tidak pernah diuji restore-nya sehingga saat darurat file backup ternyata korup. Meldir secara rutin melakukan gladi simulasi pemulihan data (backup restore drill).',
            chatTopic: 'Konsultasi Solusi Backup Otomatis & Cloud Storage'
        },
        'disaster-recovery': {
            category: 'Kelangsungan Bisnis (BCP)',
            title: 'Disaster Recovery Cepat',
            badge: 'RTO Minimal & Pemulihan',
            desc: 'Protokol pemulihan darurat sistem (Disaster Recovery) yang dirancang untuk membangkitkan kembali seluruh ekosistem aplikasi dari nol ke server baru dalam hitungan menit saat terjadi insiden katastropik.',
            impl: 'Meldir menggunakan Infrastructure as Code (IaC) dan kontainerisasi Docker/Kubernetes. Seluruh arsitektur sistem dapat direkonstruksi secara otomatis dengan target RTO (Recovery Time Objective) < 30 menit dan RPO (Recovery Point Objective) < 24 jam.',
            benefit: 'Meminimalkan downtime operasional bisnis dari yang semula bisa memakan waktu berhari-hari menjadi hitungan menit. Bisnis, kasir, dan layanan pelanggan Anda dapat segera pulih melayani transaksi.',
            tip: 'Tanpa rencana Disaster Recovery tertulis, pemulihan server yang rusak sering kali memicu kepanikan dan hilangnya data transaksi penting yang belum sempat terekam.',
            chatTopic: 'Konsultasi Perencanaan Disaster Recovery & High Availability'
        },
        'owasp-audit': {
            category: 'Keamanan Kode Aplikasi',
            title: 'Audit Celah Standar OWASP',
            badge: 'Anti SQLi, XSS & CSRF',
            desc: 'Standar audit pengujian keamanan kode mengacu pada standar global OWASP Top 10 (Open Web Application Security Project), yaitu daftar 10 celah keamanan siber paling kritis yang sering dieksploitasi peretas di dunia.',
            impl: 'Setiap modul kode diuji terhadap SQL Injection (menggunakan parameterized queries/ORM ketat), Cross-Site Scripting (output encoding DOM & Sanitizer), Cross-Site Request Forgery (anti-CSRF tokens & SameSite cookies), dan Broken Access Control.',
            benefit: 'Menutup pintu masuk utama yang paling sering dimanfaatkan hacker untuk mencuri data pengguna, memanipulasi database penjualan, atau menyusupkan script malware ke dalam sistem Anda.',
            tip: 'Audit OWASP Meldir mencakup pengujian statis (SAST) pada kode sumber dan pengujian dinamis (DAST) pada endpoint API sebelum sistem dirilis ke produksi.',
            chatTopic: 'Konsultasi Audit Keamanan OWASP Top 10'
        },
        'routine-patching': {
            category: 'Pemeliharaan Server & DevSecOps',
            title: 'Penambalan Kerentanan Rutin',
            badge: 'Patching Server Cepat',
            desc: 'Proses proaktif untuk memperbarui sistem operasi server, database engine, runtime aplikasi, dan pustaka kode eksternal (third-party dependencies) sesegera mungkin saat celah keamanan baru (CVE) ditemukan.',
            impl: 'Pemindaian kerentanan CVE mingguan dan pembaruan dependensi otomatis via tools audit keamanan (Dependabot, Snyk, Trivy). Pembaruan kernel OS dan pustaka keamanan diaplikasikan tanpa menyebabkan downtime aplikasi.',
            benefit: 'Mencegah eksploitasi celah zero-day dan celah publik yang sering diserang oleh bot peretas otomatis yang berkeliaran mencari server dengan versi software kadaluarsa di internet.',
            tip: 'Lebih dari 80% kasus peretasan web di Indonesia terjadi bukan karena hacker jenius, melainkan karena server menggunakan software lawas yang tidak pernah di-patch selama berbulan-bulan.',
            chatTopic: 'Konsultasi Layanan Pemeliharaan & Patching Rutin'
        },
        'waf-security': {
            category: 'Proteksi Perimeter Lapisan 7',
            title: 'Web Application Firewall (WAF)',
            badge: 'Filter Lalu Lintas Lapis 7',
            desc: 'Dinding pelindung cerdas di lapisan aplikasi (OSI Layer 7) yang menganalisis setiap paket permintaan data pengunjung sebelum diizinkan menyentuh server inti aplikasi bisnis Anda.',
            impl: 'Penyaringan lalu lintas menggunakan ruleset WAF mutakhir untuk mendeteksi payload berbahaya, web scraper ilegal, percobaan exploit, dan anomali perilaku trafik secara real-time di jaringan Cloudflare Enterprise / Cloud Edge.',
            benefit: 'Mengeliminasi hingga 65% beban trafik liar pada server, menolak serangan injeksi otomatis, dan melindungi aplikasi web dari eksploitasi celah baru sebelum patch server sempat dipasang.',
            tip: 'WAF berfungsi seperti pos satpam cerdas di gerbang depan gedung bisnis Anda: memeriksa identitas dan barang bawaan setiap tamu sebelum masuk ke dalam kantor.',
            chatTopic: 'Konsultasi Pemasangan Web Application Firewall'
        },
        'anti-ddos': {
            category: 'Ketahanan Jaringan & Trafik',
            title: 'Mitigasi Anti-DDoS & Bot Jahat',
            badge: 'Blokir Anomali & Scraping',
            desc: 'Sistem pertahanan terdistribusi untuk meredam serangan banjir trafik buatan (Distributed Denial of Service) dan bot berbahaya yang bertujuan membuat server kewalahan, hang, atau kehabisan memori.',
            impl: 'Penerapan rate-limiting adaptif pada endpoint login dan API publik, IP reputation scoring, serta tantangan CAPTCHA pintar berbasis JavaScript tersembunyi yang tidak mengganggu pengunjung manusia asli.',
            benefit: 'Memastikan situs dan sistem operasional bisnis Anda tetap responsif dan lancar digunakan oleh staf serta pelanggan asli, bahkan saat kompetitor curang atau peretas mencoba membanjiri server Anda.',
            tip: 'Serangan DDoS skala kecil sering kali tidak mematikan server sepenuhnya, namun membuat aplikasi sangat lambat sehingga pelanggan membatalkan pembelian mereka.',
            chatTopic: 'Konsultasi Proteksi Anti-DDoS & Rate Limiting'
        },
        'audit-trail': {
            category: 'Kepatuhan & Tata Kelola Data',
            title: 'Activity Log & Audit Trail',
            badge: 'Rekam Jejak Mutlak',
            desc: 'Sistem pencatatan kronologis yang tidak dapat diubah (immutable logs) atas setiap aktivitas, transaksi, perubahan konfigurasi, dan modifikasi data yang terjadi di dalam aplikasi.',
            impl: 'Pencatatan menyeluruh meliputi ID pengguna, timestamp UTC presisi, alamat IP asal, jenis aksi (CREATE/UPDATE/DELETE), serta snapshot data sebelum dan sesudah perubahan disimpan ke tabel audit terdedikasi.',
            benefit: 'Mencegah kecurangan internal (fraud staf), memberikan bukti sah saat terjadi sengketa transaksi, dan memenuhi persyaratan audit kepatuhan hukum transaksi digital perbankan dan ISO 27001.',
            tip: 'Audit trail yang baik harus bersifat append-only: bahkan pengguna dengan level admin tertinggi pun tidak boleh memiliki tombol untuk menghapus rekam jejak log aktivitas.',
            chatTopic: 'Konsultasi Implementasi Audit Trail & Log Transaksi'
        },
        'legal-nda': {
            category: 'Kepastian Hukum & Kerahasiaan',
            title: 'Klausul Kerahasiaan (NDA)',
            badge: 'Hukum Sah & Mengikat',
            desc: 'Perjanjian kerahasiaan formal (Non-Disclosure Agreement) yang diikat secara legal di hadapan hukum Republik Indonesia untuk melindungi ide produk, rahasia dagang, dan data pelanggan bisnis Anda.',
            impl: 'Meldir beroperasi resmi di bawah perseroan terbatas PT Melayani Digital Raya (NIB 0709260111296, SK Kemenkumham AHU-A104016.AH.01.30.Tahun 2026). Setiap kerja sama dipayungi SPK dan NDA bermaterai resmi.',
            benefit: 'Jaminan perlindungan hukum perdata dan pidana bahwa kode sumber, database rahasia, data pelanggan, dan formula bisnis Anda tidak akan pernah dibocorkan atau dijual kepada pihak ketiga.',
            tip: 'Bekerja sama dengan freelancer lepas perorangan tanpa badan hukum resmi membuat Anda sulit menuntut pertanggungjawaban hukum jika sewaktu-waktu data atau ide bisnis Anda dicuri.',
            chatTopic: 'Konsultasi Legalitas PT & Perjanjian NDA Resmi'
        },
        'copyright-ip': {
            category: 'Hak Kekayaan Intelektual (HAKI)',
            title: '100% Hak Cipta & Source Code',
            badge: 'Hak Milik Sah Klien',
            desc: 'Prinsip kepemilikan mutlak di mana seluruh kode sumber (source code), skrip arsitektur, dan hak kekayaan intelektual aplikasi diserahkan sepenuhnya menjadi milik klien tanpa biaya sewa lisensi tersembunyi.',
            impl: 'Penyerahan penuh melalui Berita Acara Serah Terima (BAST), repositori privat Git mandiri, akses penuh akun cloud provider, serta dokumentasi deployment lengkap pasca-proyek selesai.',
            benefit: 'Bebas dari vendor lock-in! Anda memiliki kendali 100% atas aset teknologi Anda. Anda bebas melanjutkan pengembangan dengan tim internal Anda sendiri kapan saja di masa depan.',
            tip: 'Banyak vendor software menerapkan model sewa lisensi tahunan sehingga jika Anda berhenti berlangganan, seluruh sistem dan data Anda disandera dan tidak bisa dipakai lagi. Meldir tidak pernah menerapkan vendor lock-in.',
            chatTopic: 'Konsultasi Kepemilikan Hak Cipta & Serah Terima Source Code'
        },
        'rbac-access': {
            category: 'Keamanan Identitas & Akses',
            title: 'Kontrol Akses Bertingkat (RBAC)',
            badge: 'Izin Granular & 2FA',
            desc: 'Pemberian hak akses secara selektif berbasis peran kerja (Role-Based Access Control) dan otentikasi ganda (Two-Factor Authentication / 2FA) untuk memastikan setiap akun hanya dapat membuka data yang diizinkan baginya.',
            impl: 'Arsitektur pemisahan hak izin granular (Owner, Manager, Staf Kasir, Auditor, Klien Luar). Sesi login dilindungi token JWT aman, token refresh rotasi, dan integrasi TOTP Authenticator (Google Authenticator).',
            benefit: 'Mencegah staf operasional melihat laporan keuangan rahasia direksi, mencegah kasir memanipulasi master stok/harga, serta mengamankan akun dari pembobolan akibat kata sandi staf yang bocor.',
            tip: 'Penerapan prinsip "Least Privilege" (hak akses minimum yang diperlukan untuk bekerja) adalah pilar pertahanan terpenting dalam mencegah kebocoran data internal perusahaan.',
            chatTopic: 'Konsultasi Konfigurasi RBAC & Keamanan 2FA'
        }
    };

    function openSecurityDetailModal(securityId) {
        if (!secModal) return;
        const data = securityStandardsData[securityId];
        if (!data) return;

        if (secModalCat) secModalCat.textContent = data.category;
        if (secModalBadge) secModalBadge.textContent = data.badge;
        if (secModalTitle) secModalTitle.textContent = data.title;
        if (secModalDesc) secModalDesc.textContent = data.desc;
        if (secModalImpl) secModalImpl.textContent = data.impl;
        if (secModalBenefit) secModalBenefit.textContent = data.benefit;
        if (secModalTip) secModalTip.textContent = data.tip;
        if (secModalChatBtn) {
            secModalChatBtn.setAttribute('data-chat-topic', data.chatTopic || 'Konsultasi Standar Keamanan');
        }

        secModal.classList.add('active');
        document.body.style.overflow = 'hidden';
    }

    function closeSecurityDetailModal() {
        if (!secModal) return;
        secModal.classList.remove('active');
        document.body.style.overflow = '';
    }

    document.querySelectorAll('[data-security-id]').forEach(el => {
        el.addEventListener('click', (e) => {
            e.preventDefault();
            const secId = el.getAttribute('data-security-id');
            openSecurityDetailModal(secId);
        });
    });

    secCloseBtns.forEach(btn => {
        btn.addEventListener('click', closeSecurityDetailModal);
    });

    if (secModal) {
        secModal.addEventListener('click', (e) => {
            if (e.target === secModal) closeSecurityDetailModal();
        });
    }

    // --- 10D. Gadget Ecosystem Platform Modals Controller ---
    const gadgetModalIds = ['gadget-modal-pc', 'gadget-modal-tablet', 'gadget-modal-smartphone', 'gadget-modal-smartwatch'];

    function closeAllGadgetModals() {
        gadgetModalIds.forEach(id => {
            const m = document.getElementById(id);
            if (m) m.classList.remove('active');
        });
        document.body.style.overflow = '';
    }

    gadgetModalIds.forEach(id => {
        const modal = document.getElementById(id);
        if (!modal) return;
        const triggers = document.querySelectorAll(`[data-modal="${id}"]`);
        triggers.forEach(btn => {
            btn.addEventListener('click', (e) => {
                e.preventDefault();
                closeAllGadgetModals();
                modal.classList.add('active');
                document.body.style.overflow = 'hidden';
            });
        });

        modal.addEventListener('click', (e) => {
            if (e.target === modal) closeAllGadgetModals();
        });
    });

    document.querySelectorAll('[data-close-gadget-modal]').forEach(btn => {
        btn.addEventListener('click', closeAllGadgetModals);
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
            if (chatWrapper && chatWrapper.classList.contains('active') && !topic) {
                closeLiveChat();
            } else {
                openLiveChat(topic);
            }
        });
    });

    function handleHashAction() {
        if (typeof checkHashRoute === 'function') checkHashRoute();
        const hash = window.location.hash.toLowerCase();
        if (hash === '#audit') {
            if (typeof openLeadModal === 'function') openLeadModal();
        } else if (hash === '#chat') {
            openLiveChat();
        }
    }
    handleHashAction();
    window.addEventListener('hashchange', handleHashAction);

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

    // --- 11. Live IDE Terminal Code Simulator (Typing Engine Infinite Loop) ---
    const terminalBody = document.getElementById('terminal-code-body');
    const terminalFileName = document.getElementById('terminal-file-name');
    const terminalMetaLang = document.getElementById('terminal-meta-lang');

    if (terminalBody) {
        const modules = [
            {
                file: 'edge_router.ts',
                lang: 'TypeScript 5.4 • Edge Runtime',
                lines: [
                    [{ t: '// Meldir Edge Gateway v2.4 (SLA 99.9%)', c: 'comment' }],
                    [{ t: 'export async function ', c: 'kw' }, { t: 'routeRequest', c: 'fn' }, { t: '(req: ', c: 'plain' }, { t: 'Request', c: 'prop' }, { t: ') {', c: 'plain' }],
                    [{ t: '  const ', c: 'kw' }, { t: 'session', c: 'id' }, { t: ' = await ', c: 'plain' }, { t: 'auth.verifyJWT', c: 'fn' }, { t: '(req);', c: 'plain' }],
                    [{ t: '  if (!session.', c: 'plain' }, { t: 'valid', c: 'prop' }, { t: ') return ', c: 'plain' }, { t: 'deny', c: 'fn' }, { t: '(401);', c: 'plain' }],
                    [{ t: '  return ', c: 'kw' }, { t: 'proxyCluster', c: 'fn' }, { t: '({ tenant: session.', c: 'plain' }, { t: 'id', c: 'prop' }, { t: ' });', c: 'plain' }],
                    [{ t: '}', c: 'plain' }]
                ]
            },
            {
                file: 'sla_watchdog.ts',
                lang: 'TypeScript 5.4 • Distributed SLA Monitor',
                lines: [
                    [{ t: '// 24/7 Node Heartbeat & Auto-Failover Monitor', c: 'comment' }],
                    [{ t: 'export async function ', c: 'kw' }, { t: 'checkClusterHealth', c: 'fn' }, { t: '(nodeId: ', c: 'plain' }, { t: 'string', c: 'prop' }, { t: ') {', c: 'plain' }],
                    [{ t: '  const ', c: 'kw' }, { t: 'probe', c: 'id' }, { t: ' = await ', c: 'plain' }, { t: 'telemetry.ping', c: 'fn' }, { t: '(nodeId, 150);', c: 'plain' }],
                    [{ t: '  if (!probe.', c: 'plain' }, { t: 'healthy', c: 'prop' }, { t: ') {', c: 'plain' }],
                    [{ t: '    await ', c: 'kw' }, { t: 'cloud.rerouteTraffic', c: 'fn' }, { t: '("sgp-1", "jkt-2");', c: 'str' }],
                    [{ t: '  }', c: 'plain' }],
                    [{ t: '}', c: 'plain' }]
                ]
            },
            {
                file: 'cluster_worker.go',
                lang: 'Go 1.22 • Multi-Thread Worker Pool',
                lines: [
                    [{ t: '// High-Throughput Job Dispatcher Pool', c: 'comment' }],
                    [{ t: 'func ', c: 'kw' }, { t: '(p *Pool) ', c: 'plain' }, { t: 'Dispatch', c: 'fn' }, { t: '(ctx ', c: 'param' }, { t: 'context.Context', c: 'prop' }, { t: ') error {', c: 'plain' }],
                    [{ t: '  select {', c: 'plain' }],
                    [{ t: '  case ', c: 'kw' }, { t: 'p.Workers <- true:', c: 'id' }],
                    [{ t: '    go ', c: 'kw' }, { t: 'p.ExecuteWithFallback', c: 'fn' }, { t: '(ctx, p.Jobs)', c: 'plain' }],
                    [{ t: '    return nil', c: 'kw' }],
                    [{ t: '  }', c: 'plain' }],
                    [{ t: '}', c: 'plain' }]
                ]
            },
            {
                file: 'auth_shield.rs',
                lang: 'Rust 1.78 • Zero-Copy Memory Guard',
                lines: [
                    [{ t: '// Zero-Allocation High-Speed Security Check', c: 'comment' }],
                    [{ t: 'pub fn ', c: 'kw' }, { t: 'validate_payload', c: 'fn' }, { t: '(bytes: &[', c: 'plain' }, { t: 'u8', c: 'prop' }, { t: ']) -> ', c: 'plain' }, { t: 'Result', c: 'prop' }, { t: '<Token, SecErr> {', c: 'plain' }],
                    [{ t: '  let ', c: 'kw' }, { t: 'hash', c: 'id' }, { t: ' = blake3::', c: 'plain' }, { t: 'hash', c: 'fn' }, { t: '(bytes);', c: 'plain' }],
                    [{ t: '  match ', c: 'kw' }, { t: 'KeyStore::get(&hash) {', c: 'plain' }],
                    [{ t: '    Some(t) => Ok(t),', c: 'plain' }],
                    [{ t: '    None => Err(SecErr::Unauthorized),', c: 'plain' }],
                    [{ t: '  }', c: 'plain' }],
                    [{ t: '}', c: 'plain' }]
                ]
            },
            {
                file: 'sqlite_sync.dart',
                lang: 'Flutter 3.22 • Offline-First Sync Engine',
                lines: [
                    [{ t: '// Realtime POS Offline Transaction Synchronization', c: 'comment' }],
                    [{ t: 'Future<', c: 'kw' }, { t: 'SyncResult', c: 'prop' }, { t: '> ', c: 'kw' }, { t: 'reconcileQueue', c: 'fn' }, { t: '() async {', c: 'plain' }],
                    [{ t: '  final ', c: 'kw' }, { t: 'pending', c: 'id' }, { t: ' = await ', c: 'plain' }, { t: 'localDb.getPendingOrders', c: 'fn' }, { t: '();', c: 'plain' }],
                    [{ t: '  for (final ', c: 'kw' }, { t: 'tx', c: 'id' }, { t: ' in pending) {', c: 'plain' }],
                    [{ t: '    await ', c: 'kw' }, { t: 'api.commitTransaction', c: 'fn' }, { t: '(tx.payload);', c: 'plain' }],
                    [{ t: '  }', c: 'plain' }],
                    [{ t: '}', c: 'plain' }]
                ]
            }
        ];

        // Start from module index 1 since module 0 (edge_router.ts) is pre-rendered in HTML
        let currentModIdx = 1;
        let lineCounter = 7;

        async function streamNextModule() {
            const mod = modules[currentModIdx];
            if (terminalFileName) terminalFileName.textContent = mod.file;
            if (terminalMetaLang) terminalMetaLang.textContent = mod.lang;

            for (let l = 0; l < mod.lines.length; l++) {
                // Remove previous active cursor and active class
                const oldCursor = terminalBody.querySelector('.code-cursor');
                if (oldCursor) oldCursor.remove();
                const oldActive = terminalBody.querySelector('.terminal-code-line.active-line');
                if (oldActive) oldActive.classList.remove('active-line');

                const lineData = mod.lines[l];
                const isComment = lineData.some(token => token.c === 'comment');

                const lineEl = document.createElement('div');
                lineEl.className = 'terminal-code-line active-line' + (isComment ? ' comment' : '');

                const numSpan = document.createElement('span');
                numSpan.className = 'line-num';
                numSpan.textContent = String(lineCounter);
                lineCounter++;
                if (lineCounter > 99) lineCounter = 1;
                lineEl.appendChild(numSpan);

                const contentSpan = document.createElement('span');
                contentSpan.className = 'line-content';
                lineEl.appendChild(contentSpan);

                const cursor = document.createElement('span');
                cursor.className = 'code-cursor';
                contentSpan.appendChild(cursor);

                terminalBody.appendChild(lineEl);

                // Smooth scroll to keep active bottom line in view
                terminalBody.scrollTo({ top: terminalBody.scrollHeight, behavior: 'smooth' });

                // Keep DOM buffer optimal by removing lines scrolled far out of view
                while (terminalBody.children.length > 9) {
                    terminalBody.firstElementChild.remove();
                }

                // Type each token in this line
                for (let t = 0; t < lineData.length; t++) {
                    const token = lineData[t];
                    const tokenSpan = document.createElement('span');
                    if (token.c && token.c !== 'plain') {
                        tokenSpan.className = 'code-' + token.c;
                    }
                    contentSpan.insertBefore(tokenSpan, cursor);

                    const text = token.t;
                    for (let c = 0; c < text.length; c++) {
                        tokenSpan.textContent += text[c];
                        await new Promise(r => setTimeout(r, 12));
                    }
                }

                // Natural brief pause after finishing the line
                await new Promise(r => setTimeout(r, 220));
            }

            // Completed module block: brief pause before streaming next module
            await new Promise(r => setTimeout(r, 1300));

            currentModIdx = (currentModIdx + 1) % modules.length;
            streamNextModule();
        }

        // Wait 1.8 seconds after page load before starting continuous stream
        setTimeout(() => {
            streamNextModule();
        }, 1800);
    }

});

// --- Register Service Worker dengan Auto-Update System ---
if ('serviceWorker' in navigator) {
    window.addEventListener('load', () => {
        const swUrl = window.location.origin ? (window.location.origin + '/sw.js') : '/sw.js';
        navigator.serviceWorker.register(swUrl)
            .then(reg => {
                // Selalu cek versi terbaru dari server saat halaman dimuat
                reg.update();

                reg.addEventListener('updatefound', () => {
                    const newWorker = reg.installing;
                    if (newWorker) {
                        newWorker.addEventListener('statechange', () => {
                            if (newWorker.state === 'installed' && navigator.serviceWorker.controller) {
                                console.log('[Meldir PWA] Versi baru tersedia! Memperbarui otomatis...');
                            }
                        });
                    }
                });
            })
            .catch(err => console.log('Service Worker registration failed: ', err));

        // Jika service worker baru mengambil alih kendali, pastikan halaman direfresh otomatis sekali
        let refreshing = false;
        navigator.serviceWorker.addEventListener('controllerchange', () => {
            if (!refreshing) {
                refreshing = true;
                window.location.reload();
            }
        });
    });
}

