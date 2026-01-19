import React, { useEffect, useRef } from 'react';
import * as Cesium from 'cesium';
import { Spin } from 'antd';

// Set Cesium Ion access token (replace with your own)
// Get your token at: https://cesium.com/ion/tokens
const CESIUM_ION_TOKEN = 'your_ion_token_here';

// Map tile providers configuration
export interface MapProvider {
  id: string;
  name: string;
  provider: any;
}

export const MAP_PROVIDERS: MapProvider[] = [
  {
    id: 'osm',
    name: 'OpenStreetMap',
    provider: new Cesium.OpenStreetMapImageryProvider({
      url: 'https://tile.openstreetmap.org/',
    }),
  },
  {
    id: 'arcgis',
    name: 'ArcGIS World Imagery',
    provider: new Cesium.ArcGisMapServerImageryProvider({
      url: 'https://services.arcgisonline.com/ArcGIS/rest/services/World_Imagery/MapServer',
    }),
  },
  {
    id: 'gaode',
    name: 'Gaode Maps (Chinese)',
    provider: new Cesium.UrlTemplateImageryProvider({
      url: 'http://webrd02.is.autonavi.com/appmaptile?lang=en&size=1&scale=1&style=7&x={x}&y={y}&z={z}',
    }),
  },
];

interface CesiumViewerProps {
  id?: string;
  className?: string;
  onViewerCreated?: (viewer: Cesium.Viewer) => void;
  initialCamera?: {
    destination: Cesium.Cartesian3;
    orientation: {
      heading: number;
      pitch: number;
      roll: number;
    };
  };
  mapProvider?: string;
  enableTerrain?: boolean;
  ionAccessToken?: string;
}

const CesiumViewer: React.FC<CesiumViewerProps> = ({
  id = 'cesium-container',
  className,
  onViewerCreated,
  initialCamera,
  mapProvider = 'osm',
  enableTerrain = false,
  ionAccessToken = CESIUM_ION_TOKEN,
}) => {
  const containerRef = useRef<HTMLDivElement>(null);
  const viewerRef = useRef<Cesium.Viewer | null>(null);

  useEffect(() => {
    if (!containerRef.current) return;

    // Set Cesium Ion access token
    if (ionAccessToken && ionAccessToken !== 'your_ion_token_here') {
      Cesium.Ion.defaultAccessToken = ionAccessToken;
    }

    // Find selected map provider
    const selectedProvider = MAP_PROVIDERS.find((p) => p.id === mapProvider);

    // Create viewer with minimal UI
    const viewer = new Cesium.Viewer(containerRef.current, {
      animation: false,
      timeline: false,
      baseLayerPicker: true,
      sceneModePicker: true,
      geocoder: false,
      homeButton: true,
      infoBox: true,
      scene3DOnly: false,
      navigationHelpButton: false,
      selectionIndicator: true,
      fullscreenButton: true,
      vrButton: false,
      baseLayer: selectedProvider?.provider,
    });

    viewerRef.current = viewer;

    // Set initial camera position if provided
    if (initialCamera) {
      viewer.camera.setView(initialCamera);
    }

    // Enable lighting based on sun position
    viewer.scene.globe.enableLighting = true;

    // Enable terrain if requested (requires Ion token)
    if (enableTerrain && ionAccessToken && ionAccessToken !== 'your_ion_token_here') {
      viewer.terrainProvider = Cesium.createWorldTerrain({
        requestWaterMask: true,
        requestVertexNormals: true,
      });
    }

    // Set Cesium base URL for static assets
    (window as any).CESIUM_BASE_URL = '/cesium/';

    // Callback
    if (onViewerCreated) {
      onViewerCreated(viewer);
    }

    // Cleanup
    return () => {
      if (viewerRef.current) {
        viewerRef.current.destroy();
        viewerRef.current = null;
      }
    };
  }, [mapProvider, enableTerrain, ionAccessToken]);

  return (
    <div
      ref={containerRef}
      id={id}
      className={className}
      style={{ width: '100%', height: '100%' }}
    />
  );
};

export default CesiumViewer;
