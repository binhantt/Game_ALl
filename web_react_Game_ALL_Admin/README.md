# Game Admin Dashboard

A modern, beautiful admin dashboard for game management, built with **React**, **TypeScript**, **Vite**, and **Ant Design**.

## Features

- ⚡️ Fast development with Vite
- 🛠️ Type safety with TypeScript
- 🎨 Beautiful UI with Ant Design
- 🧩 Modular layouts: MainLayout, AuthLayout, Sidebar, etc.
- 🔒 Authentication context (AuthContext)
- 📦 Clean folder structure for easy scaling
- 📱 Responsive design, works on all devices

## Project Structure

```
src/
  components/
    iu/
      CustomInput.tsx
  context/
    AuthContext.tsx
  layout/
    MainLayout.tsx
    SiderBar.tsx
  page/
    Home.tsx
    LoginForm.tsx
  services/
    authService.ts
  store/
    ...
  type/
    auth.ts
  App.tsx
  main.tsx
public/
  ...
```

---

Bạn chỉ cần copy/paste vào file `README.md` của dự án.  
Nếu muốn hướng dẫn chi tiết hơn về từng component, hoặc bổ sung phần API/backend, hãy nói nhé!

## Getting Started

### 1. Install dependencies

```bash
pnpm install
# or
npm install
# or
yarn install
```

### 2. Start the development server

```bash
pnpm dev
# or
npm run dev
# or
yarn dev
```

### 3. Open your browser

Go to [http://localhost:5173](http://localhost:5173) (or the port shown in your terminal).

---

## Customization

- **UI Library:** Uses [Ant Design](https://ant.design/). You can customize theme/colors in `main.tsx` or override styles as needed.
- **Sidebar/Menu:** Edit `src/layout/SiderBar.tsx` to add/remove menu items or change navigation.
- **Layouts:** Use `MainLayout` for pages with sidebar, or create more layouts in `src/layout/`.
- **Authentication:** Logic is in `src/context/AuthContext.tsx` and `src/services/authService.ts`.

---

## Linting & Formatting

- ESLint is set up for TypeScript and React.
- You can expand rules in `eslint.config.js` as your team/project grows.

---

## Build for Production

```bash
pnpm build
# or
npm run build
# or
yarn build
```

---

## Contributing

1. Fork this repo
2. Create a new branch (`git checkout -b feature/your-feature`)
3. Commit your changes
4. Push to your fork and submit a Pull Request

---

## License

MIT

---

**Made with ❤️ by [Your Name]**
