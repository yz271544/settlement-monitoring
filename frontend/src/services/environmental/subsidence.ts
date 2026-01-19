import { request } from 'umi';

export interface SubsidenceData {
  id: string;
  location: {
    longitude: number;
    latitude: number;
    altitude: number;
  };
  subsidence_rate: number;
  timestamp: string;
}

export async function fetchSubsidenceData(params?: {
  region?: string;
  start_date?: string;
  end_date?: string;
}) {
  return request<API.ResponseResult<SubsidenceData[]>>('/api/v1/environmental/subsidence', {
    method: 'GET',
    params,
  });
}

export async function fetchSubsidenceHeatmap(params?: {
  region?: string;
  date?: string;
}) {
  return request<Blob>('/api/v1/environmental/subsidence/heatmap', {
    method: 'GET',
    params,
    responseType: 'blob',
  });
}

export async function createSubsidenceRecord(data: SubsidenceData) {
  return request<API.ResponseResult<SubsidenceData>>('/api/v1/environmental/subsidence', {
    method: 'POST',
    data,
  });
}

export async function updateSubsidenceRecord(id: string, data: Partial<SubsidenceData>) {
  return request<API.ResponseResult<SubsidenceData>>(`/api/v1/environmental/subsidence/${id}`, {
    method: 'PUT',
    data,
  });
}

export async function deleteSubsidenceRecord(id: string) {
  return request<API.ResponseResult<void>>(`/api/v1/environmental/subsidence/${id}`, {
    method: 'DELETE',
  });
}
