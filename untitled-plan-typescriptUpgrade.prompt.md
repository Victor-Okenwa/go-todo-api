## Plan: Upgrade front-end from JavaScript to TypeScript

TL;DR - Migrate the Vite React client from JSX to TSX, add TypeScript tooling, update ESLint and imports, and keep the server unchanged. The client is a bun workspace.

**Steps**
1. Confirm scope: client-only migration.
2. Add TypeScript dev dependencies in `client/package.json`.
3. Create `client/tsconfig.json` with JSX and React type support.
4. Rename `client/src/App.jsx`, `client/src/main.jsx`, and `client/src/components/Todo.jsx` to `.tsx`.
5. Update imports to use `.tsx` entry paths if needed.
6. Convert the components to typed TSX:
   - type `useState` for `task` and `items`
   - add any prop or return type annotations
7. Update `client/eslint.config.js` to handle `.ts`/`.tsx` and optionally add `@typescript-eslint`.
8. Add any Vite TypeScript support stubs if required.

**Relevant files**
- `client/package.json`
- `client/tsconfig.json`
- `client/vite.config.js`
- `client/eslint.config.js`
- `client/src/App.jsx`
- `client/src/main.jsx`
- `client/src/components/Todo.jsx`

**Verification**
1. Run `bun install` in `client`.
2. Run `bun run lint` in `client`.
3. Run `bun run dev` and confirm the app loads.
4. Verify TS compile success with Vite.

**Decisions**
- Focus on client-side migration only, preserving Go server code as-is.
- Use React 19-compatible TypeScript setup.
