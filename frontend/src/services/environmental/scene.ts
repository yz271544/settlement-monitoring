import { request } from 'umi';

export interface SceneConfig {
  id: string;
  name: string;
  description: string;
  camera: {
    longitude: number;
    latitude: number;
    height: number;
    heading: number;
    pitch: number;
    roll: number;
  };
  layers: string[];
}

export async function fetchScenes() {
  return request<API.ResponseResult<SceneConfig[]>>('/api/v1/environmental/scenes', {
    method: 'GET',
  });
}

export async function fetchScene(id: string) {
  return request<API.ResponseResult<SceneConfig>>(`/api/v1/environmental/scenes/${id}`, {
    method: 'GET',
  });
}

export async function createScene(scene: Partial<SceneConfig>) {
  return request<API.ResponseResult<SceneConfig>>('/api/v1/environmental/scenes', {
    method: 'POST',
    data: scene,
  });
}

export async function updateScene(id: string, scene: Partial<SceneConfig>) {
  return request<API.ResponseResult<SceneConfig>>(`/api/v1/environmental/scenes/${id}`, {
    method: 'PUT',
    data: scene,
  });
}

export async function deleteScene(id: string) {
  return request<API.ResponseResult<void>>(`/api/v1/environmental/scenes/${id}`, {
    method: 'DELETE',
  });
}
