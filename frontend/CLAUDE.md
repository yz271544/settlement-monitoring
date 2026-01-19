# CLAUDE.md

This file provides guidance to Claude Code (claude.ai/code) when working with code in this repository.

## Project Overview

GIN-Admin-Frontend is a React-based admin dashboard frontend for the [gin-admin](https://github.com/LyricTian/gin-admin) backend. It uses Ant Design Pro components and follows an RBAC (Role-Based Access Control) architecture for permission management.

## Technology Stack

- **Framework**: UmiJS 3.5 (React framework with routing, state management, and build tools)
- **UI Library**: Ant Design 4.20 + @ant-design/pro-components 1.1
- **State Management**: DVA (included with UmiJS)
- **Styling**: Less
- **TypeScript**: 4.5
- **Node**: >=12.0.0 (tested with v16.20.2)

## Common Commands

### Development
```bash
# Start development server (dev environment, no mock)
npm run start:dev

# Start with specific environment
npm run start:test    # test environment
npm run start:pre     # pre-production environment

# Start without mock data
npm run start:no-mock
```

### Build
```bash
# Production build
npm run build

# Build with bundle analysis
npm run analyze
```

### Code Quality
```bash
# Run all linters and type check
npm run lint

# Auto-fix lint issues
npm run lint:fix

# Type check only
npm run tsc
```

### Testing
```bash
# Run unit tests
npm test

# Run component tests
npm run test:component

# Run E2E tests (Playwright)
npm run test:e2e

# Run Playwright tests directly
npm run playwright
```

### Troubleshooting Development Issues
```bash
# Clean build artifacts and reinstall (from README dev develop section)
rm -fr src/.umi
rm -fr node_modules
npm install
npm run lint:fix
npm run lint:fix
npm run dev
```

## Architecture

### Project Structure

```
src/
├── app.tsx              # Runtime configuration: initial state, request interceptors, layout
├── access.ts            # Access control (permission checking)
├── global.tsx           # Global app initialization
├── global.less          # Global styles
├── components/          # Reusable components (Footer, RightContent, etc.)
├── pages/               # Route components organized by feature
│   ├── user/           # User pages (Login, Profile)
│   └── system/         # System management pages (User, Role, Menu, Logger)
├── services/           # API service layer (auto-generated from OpenAPI + custom)
│   └── system/         # System-related API calls
├── utils/              # Utility functions
│   └── auth.ts         # Token/session management utilities
└── locales/            # i18n translation files
```

### Configuration

- **config/config.ts**: Main UmiJS configuration (plugins, routes, theme, proxy)
- **config/routes.ts**: Route definitions with codes for RBAC
- **config/proxy.ts**: API proxy configuration (dev/test environments)
- **config/defaultSettings.ts**: Layout/theme defaults
- **config/config.dev.ts**: Dev-specific plugins (react-dev-inspector)

### Key Architecture Patterns

#### 1. UmiJS Plugin Architecture
The app uses several UmiJS plugins configured in `config/config.ts`:
- `@umijs/preset-ant-design-pro`: ProComponents integration
- `plugin-layout`: Automatic layout generation
- `plugin-initial-state`: Global state management
- `plugin-access`: Permission/access control
- `plugin-locale`: Internationalization
- `plugin-model`: Alternative state management (hooks-based)
- `plugin-openapi`: Auto-generates TypeScript API services from OpenAPI spec

#### 2. Route-Based Access Control
Routes in `config/routes.ts` include a `code` field that maps to backend permissions:
```typescript
{
  path: '/system/user',
  code: 'user',           // Used for permission checking
  name: 'user',
  component: './system/User',
}
```

The `initialState.routePathCodeMap` maps route paths to permission codes, and `initialState.flatMenus` contains the user's accessible menus from the backend.

#### 3. Authentication Flow
- Tokens stored in `sessionStorage` (managed by `src/utils/auth.ts`)
- Request interceptor in `src/app.tsx` adds `Authorization: Bearer <token>` header
- API base URL configured in `src/services/index.ts` (empty = use proxy)
- Backend API: `http://localhost:8040` (proxy config)

#### 4. OpenAPI Integration
Services are partially auto-generated from `config/openapi.json`. To regenerate:
```bash
npm run openapi
```
Custom services in `src/services/system/` extend the auto-generated ones.

#### 5. Initial State Management
`src/app.tsx` exports `getInitialState()` which:
- Fetches current user (`/api/v1/current/user`)
- Fetches user menus (`/api/v1/current/menus`)
- Builds route-to-permission mapping
- Redirects to login on auth failure

#### 6. Layout Configuration
The layout plugin generates a sidebar based on user permissions. `app.tsx` exports `layout()` which:
- Filters routes based on `flatMenus` from backend
- Provides `RightContent` (user dropdown) and `Footer` components
- Adds dev-only links (OpenAPI, Components docs)

## Environment Variables

- `REACT_APP_ENV`: Environment (dev, test, pre) - determines proxy config
- `UMI_ENV`: UmiJS environment
- `MOCK`: Set to 'none' to disable mock data
- `ANALYZE`: Set to '1' for bundle analysis
- `UMI_UI`: Set to 'none' to disable UmiJS UI

## API Development

### Adding New API Endpoints
1. Add endpoint to backend OpenAPI spec
2. Run `npm run openapi` to regenerate types/services
3. Or manually add to `src/services/[module]/[name].ts`:
```typescript
import { request } from 'umi';

export async function myApi(params: API.MyParams) {
  return request<API.ResponseResult<API.MyData>>('/api/v1/my-endpoint', {
    method: 'POST',
    data: params,
  });
}
```

### TypeScript Types
API types defined in `src/services/typings.d.ts` and module-specific `typings.d.ts` files.

## Styling

- Global styles: `src/global.less`
- Component-level Less files co-located with components
- Ant Design theme customization in `config/config.ts` under `theme`
- Uses CSS variables (`'root-entry-name': 'variable'`) for runtime theme changes

## Permission System

The app implements a hierarchical permission system:
1. Routes have `code` fields (e.g., `'system.user'`)
2. Backend returns user's accessible menus with codes
3. Layout filters routes based on user permissions
4. Frontend `access.ts` can define permission check functions (currently empty)

## Notes

- `src/.umi/` contains UmiJS-generated files (do not edit manually)
- When experiencing build issues, delete `src/.umi` and `node_modules`, then reinstall
- The app uses `sessionStorage` for auth tokens (cleared on browser close)
- Proxy only works in development - production builds require CORS or same-origin API
