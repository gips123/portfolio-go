# Prompt untuk Frontend Dashboard Admin - Portfolio Management

## 🎯 Tugas
Buat **Frontend Dashboard Admin** menggunakan **Next.js** (atau React) untuk mengelola data portfolio melalui website. Dashboard ini akan terhubung ke backend API Golang yang sudah ada di `http://localhost:8080`.

**Important:** Dashboard ini adalah admin panel untuk CRUD operations (Create, Read, Update, Delete) semua data portfolio.

---

## 📋 Requirements

### Tech Stack
- **Framework:** Next.js 14+ (App Router) atau React dengan Vite
- **Styling:** Tailwind CSS
- **HTTP Client:** Axios atau Fetch API
- **Form Handling:** React Hook Form + Zod (untuk validation)
- **UI Components:** Shadcn/ui atau Headless UI (optional)
- **State Management:** React Query / TanStack Query (untuk data fetching)
- **Icons:** Lucide React atau Heroicons

---

## 🗄️ Backend API Endpoints

Base URL: `http://localhost:8080/api`

### Projects
- `GET /api/projects` - Get all projects
- `GET /api/projects?category=frontend` - Filter by category
- `GET /api/projects/:id` - Get project by ID
- `POST /api/projects` - Create project
- `PUT /api/projects/:id` - Update project
- `DELETE /api/projects/:id` - Delete project
- `GET /api/projects/categories` - Get categories
- `GET /api/projects/:id/images` - Get project images
- `POST /api/projects/:id/images` - Add project image
- `PUT /api/projects/:id/images/:imageId` - Update project image
- `DELETE /api/projects/:id/images/:imageId` - Delete project image

### About Cards
- `GET /api/about` - Get all about cards
- `GET /api/about/:id` - Get about card by ID
- `POST /api/about` - Create about card
- `PUT /api/about/:id` - Update about card
- `DELETE /api/about/:id` - Delete about card
- `GET /api/about/sidebar` - Get sidebar buttons

### Skills
- `GET /api/skills` - Get all skill categories
- `POST /api/skills` - Create skill category
- `PUT /api/skills/:id` - Update skill category
- `DELETE /api/skills/:id` - Delete skill category

### Contact
- `GET /api/contact` - Get contact data
- `POST /api/contact` - Create/Update contact data
- `PUT /api/contact` - Update contact data
- `DELETE /api/contact` - Delete contact data

---

## 🎨 UI/UX Requirements

### Layout Structure
```
┌─────────────────────────────────────────┐
│  Header: Portfolio Admin Dashboard      │
├──────────┬──────────────────────────────┤
│          │                              │
│ Sidebar  │   Main Content Area          │
│          │                              │
│ - Projects│   - Data Tables              │
│ - About  │   - Forms                    │
│ - Skills │   - Cards/List View          │
│ - Contact│   - Modals                   │
│          │                              │
└──────────┴──────────────────────────────┘
```

### Pages Needed

1. **Dashboard Home** (`/`)
   - Statistics cards (Total Projects, About Cards, Skills, Contact)
   - Recent activity
   - Quick actions

2. **Projects Management** (`/projects`)
   - List view dengan table
   - Filter by category
   - Search functionality
   - Create/Edit/Delete buttons
   - View project images

3. **Project Detail/Edit** (`/projects/:id`)
   - Form untuk edit project
   - Image gallery management
   - Add/Remove images

4. **Create Project** (`/projects/new`)
   - Form untuk create new project
   - All required fields

5. **About Cards Management** (`/about`)
   - List all about cards
   - Create/Edit/Delete
   - JSONB content editor (rich text atau JSON editor)

6. **Skills Management** (`/skills`)
   - List all skill categories
   - Create/Edit/Delete
   - Skills array editor (add/remove skills)

7. **Contact Management** (`/contact`)
   - Edit contact data
   - Contact info array editor
   - Social links array editor

---

## 📝 Form Fields

### Project Form
```typescript
{
  title: string (required)
  description: string (required, textarea)
  techStack: string[] (array, required)
  imageTitle: string
  imageDescription: string (textarea)
  imageUrl: string (required)
  buttonText: string
  detailUrl: string
  category: "frontend" | "backend" | "fullstack" | "uiux" | "mobile" | "desain-grafis" (required)
}
```

### Project Image Form
```typescript
{
  imageUrl: string (required)
  order: number (default: 0)
}
```

### About Card Form
```typescript
{
  id: string (required, unique: "about-me" | "aspirations" | "life-goals" | "hobbies" | "motto")
  title: string (required)
  icon: string
  content: {
    paragraphs: Array<{text: string, type: "highlight" | "normal"}>
    hobbies?: Array<{title: string, desc: string}>
    quote?: string
  } (JSONB)
}
```

### Skill Category Form
```typescript
{
  title: string (required)
  description: string
  icon: string
  skills: Array<{
    name: string (required)
    percentage: number (0-100, required)
    icon: string
  }>
}
```

### Contact Form
```typescript
{
  title: string
  description: string
  contactInfo: Array<{
    icon: string
    label: string
    value: string
    link: string
  }>
  socialLinks: Array<{
    name: string
    icon: string
    url: string
    color: string
  }>
}
```

---

## ✨ Features Required

### 1. Data Display
- ✅ Table view dengan pagination
- ✅ Search/filter functionality
- ✅ Sort by columns
- ✅ Responsive design (mobile-friendly)

### 2. Forms
- ✅ Form validation (client-side)
- ✅ Error handling & display
- ✅ Success notifications
- ✅ Loading states
- ✅ Auto-save draft (optional)

### 3. Image Management
- ✅ Image URL input
- ✅ Image preview
- ✅ Drag & drop untuk reorder images (optional)
- ✅ Image upload to cloud storage (optional, bisa pakai URL dulu)

### 4. JSONB Editors
- ✅ Dynamic array editors (add/remove items)
- ✅ Nested object editors
- ✅ JSON preview/validation

### 5. User Experience
- ✅ Toast notifications (success/error)
- ✅ Confirmation dialogs untuk delete
- ✅ Loading spinners
- ✅ Empty states
- ✅ Error boundaries

---

## 🎨 Design Guidelines

### Color Scheme
- Primary: Blue (#3B82F6)
- Success: Green (#10B981)
- Danger: Red (#EF4444)
- Warning: Yellow (#F59E0B)
- Background: Light gray (#F9FAFB)
- Text: Dark gray (#111827)

### Typography
- Headings: Bold, 24px-32px
- Body: Regular, 16px
- Small text: 14px

### Components Style
- Cards dengan shadow
- Rounded corners (8px)
- Consistent spacing (16px, 24px, 32px)
- Hover effects
- Smooth transitions

---

## 🔧 Technical Requirements

### API Client Setup
```typescript
// lib/api.ts
const API_BASE_URL = 'http://localhost:8080/api';

// Setup axios instance dengan interceptors
// Handle errors globally
// Add loading states
```

### State Management
- Use React Query untuk server state
- Use React Hook Form untuk form state
- Use Zustand/Jotai untuk global UI state (optional)

### Error Handling
- Global error handler
- API error messages display
- Network error handling
- Validation error display

### TypeScript
- Type semua API responses
- Type semua form data
- Type semua models

---

## 📦 Example Component Structure

```
src/
├── app/ (or pages/)
│   ├── layout.tsx
│   ├── page.tsx (Dashboard)
│   ├── projects/
│   │   ├── page.tsx (List)
│   │   ├── new/
│   │   │   └── page.tsx
│   │   └── [id]/
│   │       └── page.tsx (Edit)
│   ├── about/
│   │   └── page.tsx
│   ├── skills/
│   │   └── page.tsx
│   └── contact/
│       └── page.tsx
├── components/
│   ├── layout/
│   │   ├── Sidebar.tsx
│   │   ├── Header.tsx
│   │   └── Layout.tsx
│   ├── projects/
│   │   ├── ProjectTable.tsx
│   │   ├── ProjectForm.tsx
│   │   └── ImageGallery.tsx
│   ├── about/
│   │   ├── AboutCardList.tsx
│   │   └── AboutCardForm.tsx
│   ├── skills/
│   │   ├── SkillCategoryList.tsx
│   │   └── SkillCategoryForm.tsx
│   ├── contact/
│   │   └── ContactForm.tsx
│   └── ui/
│       ├── Button.tsx
│       ├── Input.tsx
│       ├── Modal.tsx
│       ├── Table.tsx
│       └── Toast.tsx
├── lib/
│   ├── api.ts
│   ├── types.ts
│   └── utils.ts
└── hooks/
    ├── useProjects.ts
    ├── useAbout.ts
    ├── useSkills.ts
    └── useContact.ts
```

---

## 🚀 Implementation Priority

### Phase 1 (Core Features)
1. ✅ Setup project & routing
2. ✅ Layout dengan Sidebar & Header
3. ✅ Projects list & create
4. ✅ Projects edit & delete

### Phase 2 (Complete CRUD)
5. ✅ About cards CRUD
6. ✅ Skills CRUD
7. ✅ Contact CRUD

### Phase 3 (Enhancements)
8. ✅ Image management untuk projects
9. ✅ JSONB editors (rich editing)
10. ✅ Search & filter
11. ✅ Dashboard statistics

---

## 📝 Example API Response Format

All responses follow this format:

**Success:**
```json
{
  "success": true,
  "data": { ... }
}
```

**Error:**
```json
{
  "success": false,
  "error": "Error message"
}
```

---

## 🎯 Success Criteria

Dashboard harus memiliki:
- ✅ All CRUD operations working
- ✅ Form validation
- ✅ Error handling
- ✅ Loading states
- ✅ Responsive design
- ✅ Clean, modern UI
- ✅ Good UX (intuitive navigation)

---

## 📌 Important Notes

1. **CORS:** Backend sudah dikonfigurasi untuk allow `http://localhost:3000`
2. **Environment:** Setup `.env.local` untuk API base URL
3. **Type Safety:** Gunakan TypeScript untuk type safety
4. **Error Messages:** Display user-friendly error messages
5. **Confirmation:** Always confirm before delete operations
6. **Validation:** Validate semua input sebelum submit

---

## 🔗 Additional Resources

- Backend API Documentation: `docs/CRUD_GUIDE.md`
- Postman Collection: `postman/Portfolio_API_CRUD.postman_collection.json`
- API Base URL: `http://localhost:8080`

---

**Silakan implementasikan frontend dashboard sesuai dengan requirements di atas. Gunakan best practices untuk React/Next.js development dan pastikan semua fitur berfungsi dengan baik.**

