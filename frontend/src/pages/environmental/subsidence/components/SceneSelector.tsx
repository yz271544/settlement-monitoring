import React from 'react';
import { Select, Space, Card } from 'antd';
import { EnvironmentOutlined } from '@ant-design/icons';

const { Option } = Select;

export interface Scene {
  id: string;
  name: string;
  camera: {
    destination: {
      x: number;
      y: number;
      z: number;
    };
    orientation: {
      heading: number;
      pitch: number;
      roll: number;
    };
  };
}

interface SceneSelectorProps {
  scenes: Scene[];
  currentScene: string;
  onSceneChange: (sceneId: string) => void;
  style?: React.CSSProperties;
}

const SceneSelector: React.FC<SceneSelectorProps> = ({
  scenes,
  currentScene,
  onSceneChange,
  style,
}) => {
  return (
    <Card
      size="small"
      style={{
        position: 'absolute',
        top: '20px',
        right: '20px',
        zIndex: 1000,
        minWidth: '200px',
        backgroundColor: 'rgba(255, 255, 255, 0.9)',
        backdropFilter: 'blur(10px)',
        ...style,
      }}
      bodyStyle={{ padding: '12px' }}
    >
      <Space direction="vertical" style={{ width: '100%' }} size="small">
        <Space>
          <EnvironmentOutlined style={{ color: '#1890ff' }} />
          <span style={{ fontWeight: 500, fontSize: '14px' }}>Scene View</span>
        </Space>
        <Select
          value={currentScene}
          onChange={onSceneChange}
          style={{ width: '100%' }}
          size="small"
        >
          {scenes.map((scene) => (
            <Option key={scene.id} value={scene.id}>
              {scene.name}
            </Option>
          ))}
        </Select>
      </Space>
    </Card>
  );
};

export default SceneSelector;
