# Frontend Dashboard - Quick Start Guide

Panduan cepat untuk setup dan menggunakan Frontend Dashboard Admin.

## 🚀 Quick Setup

### 1. Install Dependencies

```bash
# Jika menggunakan Next.js
npx create-next-app@latest portfolio-admin --typescript --tailwind --app

# Atau jika menggunakan Vite + React
npm create vite@latest portfolio-admin -- --template react-ts
```

### 2. Install Required Packages

```bash
# HTTP Client
npm install axios

# Form Handling
npm install react-hook-form @hookform/resolvers zod

# Data Fetching
npm install @tanstack/react-query

# UI Components (Optional - Shadcn/ui)
npx shadcn-ui@latest init

# Icons
npm install lucide-react
```

### 3. Setup Environment

Buat file `.env.local`:

```env
NEXT_PUBLIC_API_URL=http://localhost:8080/api
```

### 4. Setup API Client

Buat file `lib/api.ts`:

```typescript
import axios from 'axios';

const api = axios.create({
  baseURL: process.env.NEXT_PUBLIC_API_URL || 'http://localhost:8080/api',
  headers: {
    'Content-Type': 'application/json',
  },
});

// Response interceptor
api.interceptors.response.use(
  (response) => response.data,
  (error) => {
    const message = error.response?.data?.error || error.message;
    return Promise.reject(new Error(message));
  }
);

export default api;
```

### 5. Setup React Query

Wrap app dengan QueryClient:

```typescript
// app/layout.tsx atau _app.tsx
import { QueryClient, QueryClientProvider } from '@tanstack/react-query';

const queryClient = new QueryClient();

export default function RootLayout({ children }) {
  return (
    <QueryClientProvider client={queryClient}>
      {children}
    </QueryClientProvider>
  );
}
```

## 📝 Example: Projects List Page

```typescript
// app/projects/page.tsx
'use client';

import { useQuery, useMutation, useQueryClient } from '@tanstack/react-query';
import api from '@/lib/api';
import { useState } from 'react';

export default function ProjectsPage() {
  const queryClient = useQueryClient();
  const [category, setCategory] = useState('');

  // Fetch projects
  const { data, isLoading } = useQuery({
    queryKey: ['projects', category],
    queryFn: () => api.get('/projects', { params: { category } }),
  });

  // Delete project
  const deleteMutation = useMutation({
    mutationFn: (id: number) => api.delete(`/projects/${id}`),
    onSuccess: () => {
      queryClient.invalidateQueries({ queryKey: ['projects'] });
    },
  });

  if (isLoading) return <div>Loading...</div>;

  return (
    <div>
      <h1>Projects</h1>
      <select value={category} onChange={(e) => setCategory(e.target.value)}>
        <option value="">All</option>
        <option value="frontend">Frontend</option>
        <option value="backend">Backend</option>
      </select>
      
      <table>
        <thead>
          <tr>
            <th>ID</th>
            <th>Title</th>
            <th>Category</th>
            <th>Actions</th>
          </tr>
        </thead>
        <tbody>
          {data?.data?.map((project) => (
            <tr key={project.id}>
              <td>{project.id}</td>
              <td>{project.title}</td>
              <td>{project.category}</td>
              <td>
                <button onClick={() => deleteMutation.mutate(project.id)}>
                  Delete
                </button>
              </td>
            </tr>
          ))}
        </tbody>
      </table>
    </div>
  );
}
```

## 📝 Example: Create Project Form

```typescript
// app/projects/new/page.tsx
'use client';

import { useForm } from 'react-hook-form';
import { zodResolver } from '@hookform/resolvers/zod';
import { z } from 'zod';
import { useMutation } from '@tanstack/react-query';
import { useRouter } from 'next/navigation';
import api from '@/lib/api';

const projectSchema = z.object({
  title: z.string().min(1, 'Title is required'),
  description: z.string().min(1, 'Description is required'),
  techStack: z.array(z.string()),
  category: z.enum(['frontend', 'backend', 'fullstack', 'uiux', 'mobile', 'desain-grafis']),
  imageUrl: z.string().url('Invalid URL'),
  imageTitle: z.string().optional(),
  imageDescription: z.string().optional(),
  buttonText: z.string().optional(),
  detailUrl: z.string().optional(),
});

type ProjectForm = z.infer<typeof projectSchema>;

export default function NewProjectPage() {
  const router = useRouter();
  const { register, handleSubmit, formState: { errors } } = useForm<ProjectForm>({
    resolver: zodResolver(projectSchema),
  });

  const mutation = useMutation({
    mutationFn: (data: ProjectForm) => api.post('/projects', data),
    onSuccess: () => {
      router.push('/projects');
    },
  });

  const onSubmit = (data: ProjectForm) => {
    mutation.mutate(data);
  };

  return (
    <form onSubmit={handleSubmit(onSubmit)}>
      <div>
        <label>Title</label>
        <input {...register('title')} />
        {errors.title && <span>{errors.title.message}</span>}
      </div>

      <div>
        <label>Description</label>
        <textarea {...register('description')} />
        {errors.description && <span>{errors.description.message}</span>}
      </div>

      <div>
        <label>Category</label>
        <select {...register('category')}>
          <option value="frontend">Frontend</option>
          <option value="backend">Backend</option>
          <option value="fullstack">Fullstack</option>
          <option value="uiux">UI/UX</option>
          <option value="mobile">Mobile</option>
          <option value="desain-grafis">Desain Grafis</option>
        </select>
      </div>

      <div>
        <label>Image URL</label>
        <input {...register('imageUrl')} />
        {errors.imageUrl && <span>{errors.imageUrl.message}</span>}
      </div>

      <button type="submit" disabled={mutation.isPending}>
        {mutation.isPending ? 'Creating...' : 'Create Project'}
      </button>
    </form>
  );
}
```

## 🎨 UI Component Examples

### Button Component
```typescript
// components/ui/Button.tsx
interface ButtonProps {
  children: React.ReactNode;
  onClick?: () => void;
  variant?: 'primary' | 'danger' | 'secondary';
  disabled?: boolean;
}

export function Button({ children, onClick, variant = 'primary', disabled }: ButtonProps) {
  const baseStyles = 'px-4 py-2 rounded-lg font-medium transition-colors';
  const variants = {
    primary: 'bg-blue-600 text-white hover:bg-blue-700',
    danger: 'bg-red-600 text-white hover:bg-red-700',
    secondary: 'bg-gray-200 text-gray-800 hover:bg-gray-300',
  };

  return (
    <button
      className={`${baseStyles} ${variants[variant]}`}
      onClick={onClick}
      disabled={disabled}
    >
      {children}
    </button>
  );
}
```

### Toast Notification
```typescript
// hooks/useToast.ts
import { useState } from 'react';

export function useToast() {
  const [toast, setToast] = useState<{ message: string; type: 'success' | 'error' } | null>(null);

  const showToast = (message: string, type: 'success' | 'error' = 'success') => {
    setToast({ message, type });
    setTimeout(() => setToast(null), 3000);
  };

  return { toast, showToast };
}
```

## 🔗 Next Steps

1. Setup routing untuk semua pages
2. Create reusable components
3. Implement forms untuk semua entities
4. Add error handling & loading states
5. Style dengan Tailwind CSS
6. Add toast notifications
7. Test semua CRUD operations

---

**Lihat `PROMPT_FOR_FRONTEND_DASHBOARD.md` untuk requirements lengkap!**

