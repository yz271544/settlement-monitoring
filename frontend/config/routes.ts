export default [
  {
    path: '/user',
    layout: false,
    routes: [
      {
        name: 'login',
        path: '/user/login',
        component: './user/Login',
      },
      {
        component: './404',
      },
    ],
  },
  {
    path: '/home',
    code: 'home',
    name: 'home',
    icon: 'dashboard',
    component: './Welcome',
  },
  {
    path: '/system',
    code: 'system',
    name: 'system',
    icon: 'setting',
    routes: [
      {
        path: '/system',
        redirect: '/system/menu',
      },
      {
        path: '/system/menu',
        code: 'menu',
        name: 'menu',
        component: './system/Menu',
      },
      {
        path: '/system/role',
        code: 'role',
        name: 'role',
        component: './system/Role',
      },
      {
        path: '/system/user',
        code: 'user',
        name: 'user',
        component: './system/User',
      },
      {
        path: '/system/logger',
        code: 'logger',
        name: 'logger',
        component: './system/Logger',
      },
    ],
  },
  {
    path: '/account/profile',
    name: 'account.settings',
    component: './user/Profile',
    hideInMenu: true,
  },
  {
    path: '/environmental',
    code: 'environmental',
    name: 'environmental',
    icon: 'environment',
    routes: [
      {
        path: '/environmental',
        redirect: '/environmental/subsidence',
      },
      {
        path: '/environmental/subsidence',
        code: 'subsidence',
        name: 'subsidence',
        component: './environmental/subsidence',
      },
    ],
  },
  {
    path: '/',
    redirect: '/home',
  },
  {
    component: './404',
  },
];
