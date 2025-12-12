-- Seed Data untuk Portfolio Database
-- File ini berisi INSERT statements untuk semua data dari app/dummy

-- ============================================
-- 1. PROJECTS
-- ============================================
INSERT INTO projects (id, title, description, tech_stack, image_title, image_description, image_url, button_text, detail_url, category, created_at, updated_at) VALUES
(1, 'Disbursement Dashboard', 'A comprehensive disbursement dashboard system for managing fund transfers, merchant operations, and financial transactions with role-based access control.', ARRAY['React.js', 'TypeScript', 'Tailwind CSS'], 'Streamline Financial Operations', 'Manage fund transfers, merchant operations, and financial transactions with comprehensive analytics and role-based access control.', '/disbursement/1.png', 'View Dashboard', '#', 'frontend', NOW(), NOW()),
(2, 'Landing Page UONEPAY', 'A modern landing page for UONEPAY payment gateway with responsive design, interactive features, and seamless user experience.', ARRAY['React', 'Next.js', 'Tailwind CSS', 'TypeScript', 'Framer Motion'], 'Modern Payment Gateway Experience', 'Showcase payment gateway services with modern, professional interface featuring smooth animations and comprehensive service information.', '/uonepay/cover.png', 'Visit Website', 'https://uonepay.co', 'frontend', NOW(), NOW()),
(3, 'Payment Gateway & Disbursement Dashboard', 'A comprehensive payment gateway system with disbursement capabilities, supporting multiple payment channels, real-time transactions, and detailed financial reporting.', ARRAY['Next.js', 'Node.js', 'TypeScript'], 'Secure Financial Transactions', 'Handle secure financial transactions with multiple payment channel integrations, automated disbursement workflows, and comprehensive transaction monitoring.', '/payment/1.png', 'Explore System', '#', 'frontend', NOW(), NOW()),
(4, 'Gasstrip - Holiday Ticket Booking Website', 'A comprehensive holiday ticket booking platform with integrated payment gateway, featuring seamless ticket reservation, payment processing, and travel management system.', ARRAY['Next.js', 'Strapi', 'PostgreSQL', 'Payment Gateway', 'TypeScript'], 'Pesona Yang Memikat Untuk Perjalanan Anda', 'Temukan keindahan tersembunyi dari destinasi-destinasi menakjubkan dalam koleksi pilihan kami! Jangan sampai terlewatkan untuk menjelajahi keajaiban yang memikat bersama kami dan ciptakan momen tak terlupakan', '/travel/1.png', 'Paket Wisata', '#', 'fullstack', NOW(), NOW()),
(5, 'ifortepay Internship Experience', 'Software Development Intern at ifortepay, contributing to fintech solutions and payment gateway development with hands-on experience in modern web technologies.', ARRAY['React.js', 'Node.js', 'FinTech', 'API Integration', 'Next.js', 'Strapi'], 'Real-World FinTech Development', 'Gained comprehensive experience in fintech development, working on payment gateway solutions and financial technology applications with real-world exposure.', '/ifortepay/cover.png', 'View Experience', 'https://ifortepay.id', 'fullstack', NOW(), NOW()),
(6, 'Globementor - Holiday Booking Website UI/UX', 'A comprehensive travel and holiday booking website UI/UX design with intuitive room reservation system, destination discovery, and seamless booking experience for travelers.', ARRAY['Figma', 'Canva'], 'Seamless Travel Booking Experience', 'Simplify holiday planning and accommodation booking with intuitive interface for discovering destinations, comparing room options, and managing bookings.', '/globemen/cover.png', 'View Design', '#', 'uiux', NOW(), NOW()),
(7, 'Kost Finance - Student Financial Management App UI/UX', 'A mobile application UI/UX design specifically for students living in boarding houses (kost) to manage their finances, track expenses, and plan budgets effectively.', ARRAY['Figma'], 'Smart Financial Management for Students', 'Help students manage their limited budgets effectively with intuitive expense tracking, budget planning tools, and financial insights designed for student life.', '/kostfinenace/cover.png', 'View Design', '#', 'uiux', NOW(), NOW()),
(8, 'Design Portfolio 1', 'Creative design work showcasing modern visual aesthetics and innovative layout concepts.', ARRAY['Adobe Illustrator', 'Photoshop', 'Canva'], 'Modern Visual Aesthetics', 'Showcasing innovative layout concepts and creative design work with modern visual aesthetics.', '/design/1.jpg', 'View Portfolio', '#', 'desain-grafis', NOW(), NOW()),
(9, 'Design Portfolio 2', 'Professional design project with clean typography and sophisticated color palette.', ARRAY['Adobe Illustrator', 'Photoshop'], 'Professional Design Excellence', 'Professional design project featuring clean typography and sophisticated color palette for modern branding.', '/design/2.jpg', 'View Design', '#', 'desain-grafis', NOW(), NOW()),
(10, 'Design Portfolio 3', 'Bold and dynamic design featuring striking visual elements and contemporary style.', ARRAY['Adobe Illustrator', 'Photoshop', 'Canva'], 'Bold & Dynamic Design', 'Striking visual elements and contemporary style in a bold and dynamic design approach.', '/design/6.jpg', 'Explore Design', '#', 'desain-grafis', NOW(), NOW()),
(11, 'Design Portfolio 4', 'Elegant design solution with refined aesthetics and attention to detail.', ARRAY['Adobe Illustrator', 'Photoshop'], 'Elegant Design Solution', 'Refined aesthetics and attention to detail in an elegant design solution.', '/design/4.PNG', 'View Portfolio', '#', 'desain-grafis', NOW(), NOW()),
(12, 'Design Portfolio 5', 'Comprehensive design project showcasing versatility and creative problem-solving.', ARRAY['Adobe Illustrator', 'Photoshop', 'Canva'], 'Versatile Creative Solutions', 'Showcasing versatility and creative problem-solving in comprehensive design projects.', '/design/7.PNG', 'View Design', '#', 'desain-grafis', NOW(), NOW()),
(13, 'Design Portfolio 6', 'Modern design approach with innovative concepts and user-centered thinking.', ARRAY['Adobe Illustrator', 'Photoshop'], 'Innovative User-Centered Design', 'Modern design approach featuring innovative concepts and user-centered thinking.', '/design/3.PNG', 'Explore Design', '#', 'desain-grafis', NOW(), NOW()),
(14, 'Design Portfolio 7', 'Artistic design work with creative flair and expressive visual communication.', ARRAY['Adobe Illustrator', 'Photoshop', 'Canva'], 'Artistic Creative Expression', 'Creative flair and expressive visual communication in artistic design work.', '/design/5.PNG', 'View Portfolio', '#', 'desain-grafis', NOW(), NOW()),
(15, 'Design Portfolio 8', 'Professional design solution with strategic thinking and brand consistency.', ARRAY['Adobe Illustrator', 'Photoshop'], 'Strategic Brand Design', 'Strategic thinking and brand consistency in professional design solutions.', '/design/8.PNG', 'View Design', '#', 'desain-grafis', NOW(), NOW()),
(16, 'Design Portfolio 9', 'Innovative design project pushing creative boundaries and exploring new possibilities.', ARRAY['Adobe Illustrator', 'Photoshop', 'Canva'], 'Pushing Creative Boundaries', 'Exploring new possibilities and pushing creative boundaries in innovative design projects.', '/design/9.PNG', 'Explore Design', '#', 'desain-grafis', NOW(), NOW()),
(17, 'Design Portfolio 10', 'Final design showcase representing the culmination of creative vision and technical skill.', ARRAY['Adobe Illustrator', 'Photoshop'], 'Creative Vision & Technical Skill', 'The culmination of creative vision and technical skill in final design showcase.', '/design/10.PNG', 'View Portfolio', '#', 'desain-grafis', NOW(), NOW()),
(18, 'Bendaku - Mobile Application', 'A comprehensive mobile application designed for modern users with intuitive interface, seamless navigation, and engaging user experience.', ARRAY['React Native', 'TypeScript', 'Mobile Development'], 'Modern Mobile Experience', 'Experience seamless mobile interactions with intuitive design, smooth navigation, and comprehensive features designed for modern mobile users.', '/bendaku/cover.png', 'View App', '#', 'mobile', NOW(), NOW()),
(19, 'Bendaku Backend - Strapi CMS', 'A robust backend system built with Strapi CMS for managing content, APIs, and data for the Bendaku mobile application with comprehensive admin panel and API endpoints.', ARRAY['Strapi', 'Node.js', 'REST API', 'CMS'], 'Powerful Backend Infrastructure', 'Comprehensive backend solution with Strapi CMS providing flexible content management, secure API endpoints, and efficient data handling for mobile applications.', '/bendaku_backend/1.png', 'View Backend', '#', 'backend', NOW(), NOW()),
(20, 'Design Portfolio 11', 'Creative event poster design with modern aesthetic and engaging visual elements.', ARRAY['Adobe Illustrator', 'Photoshop', 'Canva'], 'Creative Event Design', 'Modern event poster showcasing creative design work with engaging visual elements and professional aesthetics.', '/design/11.jpg', 'View Design', '#', 'desain-grafis', NOW(), NOW()),
(21, 'Design Portfolio 12', 'Retro-themed event poster with vibrant colors and vintage digital camera aesthetic.', ARRAY['Adobe Illustrator', 'Photoshop'], 'Retro Event Aesthetic', 'Vibrant retro-themed event poster featuring vintage digital camera aesthetic with chromatic aberration effects and bold typography.', '/design/12.png', 'View Design', '#', 'desain-grafis', NOW(), NOW()),
(22, 'Design Portfolio 13', 'Event reschedule announcement with calendar design and professional layout.', ARRAY['Adobe Illustrator', 'Photoshop', 'Canva'], 'Event Reschedule Design', 'Professional event reschedule announcement featuring calendar design with clear date highlighting and engaging visual presentation.', '/design/13.png', 'View Design', '#', 'desain-grafis', NOW(), NOW()),
(23, 'ifortepay Internship Experience - Frontend', 'Frontend development work at ifortepay, building responsive user interfaces and interactive features for fintech solutions using modern React and Next.js technologies.', ARRAY['React.js', 'Next.js', 'TypeScript', 'Tailwind CSS', 'Framer Motion'], 'Frontend FinTech Development', 'Developed responsive and interactive frontend interfaces for payment gateway solutions with modern React ecosystem and best practices.', '/ifortepay/cover.png', 'View Experience', 'https://ifortepay.id', 'frontend', NOW(), NOW()),
(24, 'ifortepay Internship Experience - Backend', 'Backend development work at ifortepay, developing APIs, integrating payment systems, and building server-side solutions for fintech applications using Node.js and Strapi.', ARRAY['Node.js', 'Strapi', 'REST API', 'PostgreSQL', 'API Integration'], 'Backend FinTech Development', 'Built robust backend systems and APIs for payment gateway solutions, handling secure transactions and financial data processing.', '/ifortepay/cover.png', 'View Experience', 'https://ifortepay.id', 'backend', NOW(), NOW())
ON CONFLICT (id) DO UPDATE SET
    title = EXCLUDED.title,
    description = EXCLUDED.description,
    tech_stack = EXCLUDED.tech_stack,
    image_title = EXCLUDED.image_title,
    image_description = EXCLUDED.image_description,
    image_url = EXCLUDED.image_url,
    button_text = EXCLUDED.button_text,
    detail_url = EXCLUDED.detail_url,
    category = EXCLUDED.category,
    updated_at = NOW();

-- ============================================
-- 2. PROJECT IMAGES
-- ============================================
INSERT INTO project_images (project_id, image_url, "order", created_at) VALUES
-- Project 1 (Disbursement)
(1, '/disbursement/login.png', 0, NOW()),
(1, '/disbursement/1.png', 1, NOW()),
(1, '/disbursement/2.png', 2, NOW()),
(1, '/disbursement/3.png', 3, NOW()),
(1, '/disbursement/4.png', 4, NOW()),
(1, '/disbursement/5.png', 5, NOW()),
(1, '/disbursement/6.png', 6, NOW()),
(1, '/disbursement/7.png', 7, NOW()),
(1, '/disbursement/8.png', 8, NOW()),
(1, '/disbursement/9.png', 9, NOW()),
(1, '/disbursement/10.png', 10, NOW()),

-- Project 2 (UONEPAY)
(2, '/uonepay/cover.png', 0, NOW()),
(2, '/uonepay/1.png', 1, NOW()),
(2, '/uonepay/2.png', 2, NOW()),
(2, '/uonepay/3.png', 3, NOW()),
(2, '/uonepay/4.png', 4, NOW()),
(2, '/uonepay/5.png', 5, NOW()),

-- Project 3 (Payment)
(3, '/payment/1.png', 0, NOW()),
(3, '/payment/2.png', 1, NOW()),
(3, '/payment/3.png', 2, NOW()),
(3, '/payment/4.png', 3, NOW()),

-- Project 4 (Travel)
(4, '/travel/1.png', 0, NOW()),
(4, '/travel/2.png', 1, NOW()),
(4, '/travel/3.png', 2, NOW()),
(4, '/travel/4.png', 3, NOW()),
(4, '/travel/6.png', 4, NOW()),
(4, '/travel/7.png', 5, NOW()),
(4, '/travel/8.png', 6, NOW()),

-- Project 5 (ifortepay)
(5, '/ifortepay/cover.png', 0, NOW()),
(5, '/ifortepay/1.png', 1, NOW()),
(5, '/ifortepay/2.png', 2, NOW()),
(5, '/ifortepay/3.png', 3, NOW()),
(5, '/ifortepay/4.png', 4, NOW()),
(5, '/ifortepay/5.png', 5, NOW()),
(5, '/ifortepay/6.png', 6, NOW()),
(5, '/ifortepay/7.png', 7, NOW()),
(5, '/ifortepay/8.png', 8, NOW()),
(5, '/ifortepay/9.png', 9, NOW()),

-- Project 6 (Globemen)
(6, '/globemen/cover.png', 0, NOW()),
(6, '/globemen/1.png', 1, NOW()),
(6, '/globemen/2.png', 2, NOW()),
(6, '/globemen/3.png', 3, NOW()),

-- Project 7 (Kost Finance)
(7, '/kostfinenace/cover.png', 0, NOW()),
(7, '/kostfinenace/1.png', 1, NOW()),
(7, '/kostfinenace/2.png', 2, NOW()),
(7, '/kostfinenace/3.png', 3, NOW()),
(7, '/kostfinenace/4.png', 4, NOW()),
(7, '/kostfinenace/5.png', 5, NOW()),
(7, '/kostfinenace/6.png', 6, NOW()),

-- Project 18 (Bendaku)
(18, '/bendaku/cover.png', 0, NOW()),
(18, '/bendaku/login.png', 1, NOW()),
(18, '/bendaku/1.png', 2, NOW()),
(18, '/bendaku/2.png', 3, NOW()),
(18, '/bendaku/3.png', 4, NOW()),
(18, '/bendaku/4.png', 5, NOW()),
(18, '/bendaku/5.png', 6, NOW()),
(18, '/bendaku/6.png', 7, NOW()),
(18, '/bendaku/7.png', 8, NOW()),
(18, '/bendaku/8.png', 9, NOW()),
(18, '/bendaku/9.png', 10, NOW()),

-- Project 19 (Bendaku Backend)
(19, '/bendaku_backend/1.png', 0, NOW()),
(19, '/bendaku_backend/2.png', 1, NOW()),
(19, '/bendaku_backend/3.png', 2, NOW()),
(19, '/bendaku_backend/4.png', 3, NOW()),
(19, '/bendaku_backend/5.png', 4, NOW()),
(19, '/bendaku_backend/6.png', 5, NOW()),

-- Project 23 (ifortepay Frontend)
(23, '/ifortepay/cover.png', 0, NOW()),
(23, '/ifortepay/1.png', 1, NOW()),
(23, '/ifortepay/2.png', 2, NOW()),
(23, '/ifortepay/3.png', 3, NOW()),
(23, '/ifortepay/4.png', 4, NOW()),
(23, '/ifortepay/5.png', 5, NOW()),
(23, '/ifortepay/6.png', 6, NOW()),
(23, '/ifortepay/7.png', 7, NOW()),
(23, '/ifortepay/8.png', 8, NOW()),
(23, '/ifortepay/9.png', 9, NOW()),

-- Project 24 (ifortepay Backend)
(24, '/ifortepay/cover.png', 0, NOW()),
(24, '/ifortepay/1.png', 1, NOW()),
(24, '/ifortepay/2.png', 2, NOW()),
(24, '/ifortepay/3.png', 3, NOW()),
(24, '/ifortepay/4.png', 4, NOW()),
(24, '/ifortepay/5.png', 5, NOW()),
(24, '/ifortepay/6.png', 6, NOW()),
(24, '/ifortepay/7.png', 7, NOW()),
(24, '/ifortepay/8.png', 8, NOW()),
(24, '/ifortepay/9.png', 9, NOW());

-- ============================================
-- 3. ABOUT CARDS
-- ============================================
INSERT INTO about_cards (id, title, icon, content, created_at, updated_at) VALUES
('about-me', 'ABOUT ME', 'User', 
'{"paragraphs": [{"text": "Hello! I''m <span className=\"font-semibold text-white\">Ghifary Ahmad Alfirdausy</span>, a passionate Full Stack Developer dedicated to creating beautiful, functional, and user-friendly web experiences across all layers of application development.", "type": "highlight"}, {"text": "With a strong foundation in modern web technologies, I specialize in building responsive and interactive applications from frontend to backend and database management. I''m skilled in creating applications that not only look great but also provide exceptional user experiences with robust architecture and efficient data management. I''m constantly learning and exploring new technologies to stay at the forefront of full-stack development.", "type": "normal"}]}'::jsonb,
NOW(), NOW()),

('aspirations', 'ASPIRATIONS', 'Rocket',
'{"paragraphs": [{"text": "My aspiration is to become a <span className=\"font-semibold text-white\">Full Stack Developer</span> who is skilled and recognized in all aspects of technology - from frontend, backend, to database management. I want to continuously develop my abilities in building complete web applications that are not only visually beautiful but also robust, scalable, and provide exceptional user experiences.", "type": "highlight"}, {"text": "I aspire to master all layers of application development, contribute to large-scale projects that have a positive impact on many people, and become a mentor for beginner developers who want to learn and grow in the field of full-stack development.", "type": "normal"}]}'::jsonb,
NOW(), NOW()),

('life-goals', 'LIFE GOALS', 'Target',
'{"paragraphs": [{"text": "My life goal is to keep learning and growing, both professionally and personally. I want to use my abilities and knowledge to create technological solutions that can help solve real problems in everyday life.", "type": "highlight"}, {"text": "Additionally, I also want to build a stable and successful career while maintaining a balance between work and personal life. I believe that with dedication, hard work, and a continuous learning spirit, I can achieve all the goals I have set.", "type": "normal"}]}'::jsonb,
NOW(), NOW()),

('hobbies', 'HOBBIES', 'Heart',
'{"paragraphs": [{"text": "In my free time, I enjoy various activities that help me stay creative and inspired. Some of my hobbies include:", "type": "normal"}], "hobbies": [{"title": "Coding & Learning", "desc": "Exploring new technologies and building personal projects"}, {"title": "Design", "desc": "Creating designs and following the latest design trends"}, {"title": "Music", "desc": "Listening to music and discovering new sounds for inspiration and relaxation"}, {"title": "Billiard", "desc": "Playing billiard for relaxation and strategic thinking"}]}'::jsonb,
NOW(), NOW()),

('motto', 'MOTTO', 'Quote',
'{"quote": "Never stop learning, because life never stops teaching.", "paragraphs": [{"text": "This motto reminds me to always be open to new learning and never feel satisfied with the knowledge I already have. Every day is an opportunity to learn something new and become a better version of myself.", "type": "normal"}]}'::jsonb,
NOW(), NOW())
ON CONFLICT (id) DO UPDATE SET
    title = EXCLUDED.title,
    icon = EXCLUDED.icon,
    content = EXCLUDED.content,
    updated_at = NOW();

-- ============================================
-- 4. SKILL CATEGORIES
-- ============================================
INSERT INTO skill_categories (id, title, description, icon, skills, created_at, updated_at) VALUES
(1, 'Frontend Development', 'Creating beautiful, responsive, and interactive user interfaces with modern frameworks.', 'Monitor',
'[{"name": "React/Next.js", "percentage": 90, "icon": "Code"}, {"name": "TypeScript", "percentage": 88, "icon": "Code"}, {"name": "Tailwind CSS", "percentage": 90, "icon": "Code"}, {"name": "Framer Motion", "percentage": 80, "icon": "Code"}]'::jsonb,
NOW(), NOW()),

(2, 'Backend Development', 'Building robust, scalable server-side applications and APIs with modern technologies.', 'Server',
'[{"name": "Node.js", "percentage": 80, "icon": "Server"}, {"name": "Express.js", "percentage": 80, "icon": "Server"}, {"name": "MySQL", "percentage": 88, "icon": "Database"}, {"name": "MongoDB", "percentage": 70, "icon": "Database"}]'::jsonb,
NOW(), NOW()),

(3, 'UI/UX Design', 'Designing intuitive and aesthetically pleasing user experiences.', 'Palette',
'[{"name": "Figma", "percentage": 80, "icon": "PenTool"}, {"name": "Canva", "percentage": 95, "icon": "PenTool"}, {"name": "Prototyping", "percentage": 80, "icon": "PenTool"}, {"name": "User Research", "percentage": 88, "icon": "Users"}]'::jsonb,
NOW(), NOW())
ON CONFLICT (id) DO UPDATE SET
    title = EXCLUDED.title,
    description = EXCLUDED.description,
    icon = EXCLUDED.icon,
    skills = EXCLUDED.skills,
    updated_at = NOW();

-- ============================================
-- 5. CONTACT DATA
-- ============================================
INSERT INTO contact_data (id, title, description, contact_info, social_links, updated_at) VALUES
(1, 'CONTACT', 'Let''s work together! I''m always open to discussing new projects, creative ideas, or opportunities to be part of your visions.',
'[{"icon": "Mail", "label": "Email", "value": "ghifaryahmadwrap@gmail.com", "link": "mailto:ghifaryahmadwrap@gmail.com"}, {"icon": "MapPin", "label": "Location", "value": "Jakarta, Indonesia", "link": "#"}, {"icon": "Phone", "label": "Phone", "value": "+62 852-1900-8008", "link": "tel:+6285219008008"}]'::jsonb,
'[{"name": "LinkedIn", "icon": "Linkedin", "url": "https://linkedin.com/in/ghifaryahmada", "color": "hover:text-blue-400"}, {"name": "GitHub", "icon": "Github", "url": "https://github.com/gips123", "color": "hover:text-gray-300"}, {"name": "Instagram", "icon": "Instagram", "url": "https://instagram.com/ghifaryahmada", "color": "hover:text-pink-400"}]'::jsonb,
NOW())
ON CONFLICT (id) DO UPDATE SET
    title = EXCLUDED.title,
    description = EXCLUDED.description,
    contact_info = EXCLUDED.contact_info,
    social_links = EXCLUDED.social_links,
    updated_at = NOW();

