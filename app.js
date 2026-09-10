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

        // Translate WhatsApp dynamic URLs (for index.html)
        const heroWaBtns = document.querySelectorAll('.hero-btn-wa, #hero-cta-wa');
        heroWaBtns.forEach(btn => btn.setAttribute('href', waTemplates.hero[lang]));
        const socialWaBtn = document.querySelector('.social-content .btn');
        if (socialWaBtn) socialWaBtn.setAttribute('href', waTemplates.social[lang]);
        const quickWaBtn = document.querySelector('.sidebar-box.box-wa');
        if (quickWaBtn) quickWaBtn.setAttribute('href', waTemplates.quick[lang]);
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
        const isOpen = mobileToggle.classList.toggle('open');
        mobileNavOverlay.classList.toggle('open', isOpen);
        document.body.style.overflow = isOpen ? 'hidden' : '';
    }

    function closeMobileMenu() {
        mobileToggle.classList.remove('open');
        mobileNavOverlay.classList.remove('open');
        document.body.style.overflow = '';
    }

    mobileToggle.addEventListener('click', toggleMobileMenu);
    
    mobileNavLinks.forEach(link => {
        link.addEventListener('click', closeMobileMenu);
    });

    // Close mobile menu on resize if screen gets larger
    window.addEventListener('resize', () => {
        if (window.innerWidth > 768) {
            closeMobileMenu();
        }
    });


    // --- 3. Scroll Header Styles & Active Menu ---
    const header = document.getElementById('site-header');
    const navItems = document.querySelectorAll('.nav-item');
    const sections = document.querySelectorAll('section[id]');

    function handleScroll() {
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
            // Hanya aktif untuk tampilan desktop (> 1024px)
            if (window.innerWidth <= 1024) {
                quickSidebar.classList.add('collapsed');
                return;
            }

            const scrollY = window.scrollY;
            const heroHeight = heroSection ? heroSection.offsetHeight : 600;

            // Selama layar menampilkan kondisi di section atas/hero: selalu TERBUKA
            // Tertutup otomatis jika scroll ke bawah melewati hero, dan kembali terbuka jika naik ke section tersebut
            if (scrollY < heroHeight - 120) {
                quickSidebar.classList.remove('collapsed');
            } else {
                quickSidebar.classList.add('collapsed');
            }
        }

        window.addEventListener('scroll', updateSidebarByScroll, { passive: true });
        window.addEventListener('resize', updateSidebarByScroll, { passive: true });
        updateSidebarByScroll(); // Inisialisasi saat pertama dimuat
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

