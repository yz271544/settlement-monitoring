// https://umijs.org/config/
import { defineConfig } from 'umi';
import { join } from 'path';

import defaultSettings from './defaultSettings';
import proxy from './proxy';
import routes from './routes';

const { REACT_APP_ENV } = process.env;

export default defineConfig({
  hash: true,
  antd: {},
  dva: {
    hmr: true,
  },
  layout: {
    // https://umijs.org/zh-CN/plugins/plugin-layout
    locale: true,
    siderWidth: 218,
    ...defaultSettings,
  },
  // https://umijs.org/zh-CN/plugins/plugin-locale
  locale: {
    // default zh-CN
    default: 'en-US',
    antd: true,
    // default true, when it is true, will use `navigator.language` overwrite default
    baseNavigator: true,
  },
  dynamicImport: {
    loading: '@ant-design/pro-layout/es/PageLoading',
  },
  targets: {
    ie: 11,
  },
  // umi routes: https://umijs.org/docs/routing
  routes,
  access: {},
  // Theme for antd: https://ant.design/docs/react/customize-theme-cn
  theme: {
    // 如果不想要 configProvide 动态设置主题需要把这个设置为 default
    // 只有设置为 variable， 才能使用 configProvide 动态设置主色调
    // https://ant.design/docs/react/customize-theme-variable-cn
    'root-entry-name': 'variable',
  },
  // esbuild is father build tools
  // https://umijs.org/plugins/plugin-esbuild
  esbuild: {},
  title: false,
  ignoreMomentLocale: true,
  proxy: proxy[REACT_APP_ENV || 'dev'],
  manifest: {
    basePath: '/',
  },
  fastRefresh: {},
  openAPI: [
    {
      requestLibPath: "import { request } from 'umi'",
      schemaPath: join(__dirname, 'openapi.json'),
      mock: false,
    },
  ],
  nodeModulesTransform: { type: 'none' },
  mfsu: {},
  webpack5: {},
  exportStatic: {},
  // Cesium configuration
  chainWebpack(config: any) {
    // Copy Cesium static files (Workers, Assets, Widgets, ThirdParty)
    // config.plugin('copy').use('copy-webpack-plugin', [
    //   {
    //     patterns: [
    //       {
    //         from: 'node_modules/cesium/Build/Cesium/Workers',
    //         to: 'cesium/Workers',
    //       },
    //       {
    //         from: 'node_modules/cesium/Build/Cesium/ThirdParty',
    //         to: 'cesium/ThirdParty',
    //       },
    //       {
    //         from: 'node_modules/cesium/Build/Cesium/Assets',
    //         to: 'cesium/Assets',
    //       },
    //       {
    //         from: 'node_modules/cesium/Build/Cesium/Widgets',
    //         to: 'cesium/Widgets',
    //       },
    //     ],
    //   },
    // ]);
    config.plugin('copy').tap((args: any[]) => {
      const existing = (args[0] && args[0].patterns) ? args[0].patterns : [];

      existing.push(
        { from: 'node_modules/cesium/Build/Cesium/Workers', to: 'cesium/Workers' },
        { from: 'node_modules/cesium/Build/Cesium/ThirdParty', to: 'cesium/ThirdParty' },
        { from: 'node_modules/cesium/Build/Cesium/Assets', to: 'cesium/Assets' },
        { from: 'node_modules/cesium/Build/Cesium/Widgets', to: 'cesium/Widgets' },
      );

      return [{ patterns: existing, options: args[0]?.options || {} }];
    });
  },
  // Define Cesium base URL
  define: {
    CESIUM_BASE_URL: JSON.stringify('/cesium/'),
  },
  // Extra Cesium public path
  publicPath: '/',
});
