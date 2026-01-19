import React, { useState, useCallback } from 'react';
import { PageContainer } from '@ant-design/pro-components';
import { Card } from 'antd';
import { useIntl } from 'umi';
import * as Cesium from 'cesium';
import CesiumViewer, { MAP_PROVIDERS } from './components/CesiumViewer';
import type { Scene } from './components/SceneSelector';
import SceneSelector from './components/SceneSelector';
import styles from './index.less';

// Define available 3D scenes for subsidence monitoring
const SCENES: Scene[] = [
  {
    id: 'global',
    name: 'Global View',
    camera: {
      destination: Cesium.Cartesian3.fromDegrees(0, 20, 20000000),
      orientation: {
        heading: Cesium.Math.toRadians(0),
        pitch: Cesium.Math.toRadians(-90),
        roll: 0.0,
      },
    },
  },
  {
    id: 'china',
    name: 'China View',
    camera: {
      destination: Cesium.Cartesian3.fromDegrees(104.1954, 35.8617, 5000000),
      orientation: {
        heading: Cesium.Math.toRadians(0),
        pitch: Cesium.Math.toRadians(-60),
        roll: 0.0,
      },
    },
  },
  {
    id: 'beijing',
    name: 'Beijing Area',
    camera: {
      destination: Cesium.Cartesian3.fromDegrees(116.39, 39.9, 500000),
      orientation: {
        heading: Cesium.Math.toRadians(0),
        pitch: Cesium.Math.toRadians(-45),
        roll: 0.0,
      },
    },
  },
  {
    id: 'local',
    name: 'Local Monitoring',
    camera: {
      destination: Cesium.Cartesian3.fromDegrees(116.39, 39.9, 10000),
      orientation: {
        heading: Cesium.Math.toRadians(0),
        pitch: Cesium.Math.toRadians(-30),
        roll: 0.0,
      },
    },
  },
];

const SubsidenceMonitoring: React.FC = () => {
  const intl = useIntl();
  const [currentScene, setCurrentScene] = useState('global');
  const [currentMap, setCurrentMap] = useState('osm');
  const [viewer, setViewer] = useState<Cesium.Viewer | null>(null);

  const handleViewerCreated = useCallback((cesiumViewer: Cesium.Viewer) => {
    setViewer(cesiumViewer);

    // Add terrain (optional - requires Cesium Ion token)
    // Uncomment if you have a valid Ion token
    // cesiumViewer.terrainProvider = Cesium.createWorldTerrain({
    //   requestWaterMask: true,
    //   requestVertexNormals: true,
    // });

    // Load initial imagery layer
    const osmLayer = new Cesium.OpenStreetMapImageryProvider({
      url: 'https://tile.openstreetmap.org/',
    });
    cesiumViewer.imageryLayers.addImageryProvider(osmLayer);

    // Enable lighting based on sun position
    cesiumViewer.scene.globe.enableLighting = true;
  }, []);

  const handleSceneChange = useCallback(
    (sceneId: string) => {
      if (!viewer) return;

      const scene = SCENES.find((s) => s.id === sceneId);
      if (scene) {
        // Fly to the new scene
        viewer.camera.flyTo({
          destination: scene.camera.destination,
          orientation: scene.camera.orientation,
          duration: 2, // 2 seconds flight animation
        });
        setCurrentScene(sceneId);
      }
    },
    [viewer],
  );

  const handleMapChange = useCallback(
    (mapId: string) => {
      if (!viewer) return;

      const provider = MAP_PROVIDERS.find((p) => p.id === mapId);
      if (provider) {
        // Remove all existing imagery layers
        viewer.imageryLayers.removeAll();
        // Add new imagery layer
        viewer.imageryLayers.addImageryProvider(provider.provider);
        setCurrentMap(mapId);
      }
    },
    [viewer],
  );

  return (
    <PageContainer
      title={intl.formatMessage({ id: 'pages.environmental.subsidence.title', defaultMessage: 'Geological Subsidence Monitoring' })}
      content={intl.formatMessage({ id: 'pages.environmental.subsidence.description', defaultMessage: '3D visualization of ground subsidence data' })}
    >
      <Card className={styles.mapCard}>
        <div className={styles.cesiumContainer}>
          <SceneSelector
            scenes={SCENES}
            currentScene={currentScene}
            onSceneChange={handleSceneChange}
          />
          <CesiumViewer
            onViewerCreated={handleViewerCreated}
            initialCamera={SCENES[0].camera}
            mapProvider={currentMap}
            enableTerrain={false}
          />
        </div>
      </Card>
    </PageContainer>
  );
};

export default SubsidenceMonitoring;
