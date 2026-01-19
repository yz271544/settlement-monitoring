declare namespace API {
  // Environmental monitoring types
  type SubsidencePoint = {
    id: string;
    longitude: number;
    latitude: number;
    altitude: number;
    subsidence_value: number;
    measurement_date: string;
    region?: string;
  };

  type EnvironmentalScene = {
    id: string;
    name: string;
    description: string;
    bounds: {
      west: number;
      south: number;
      east: number;
      north: number;
    };
    camera_position: {
      longitude: number;
      latitude: number;
      height: number;
    };
    camera_orientation: {
      heading: number;
      pitch: number;
      roll: number;
    };
  };

  type SubsidenceHeatmapData = {
    region: string;
    date: string;
    data_url: string;
    min_value: number;
    max_value: number;
  };

  type MonitoringStation = {
    id: string;
    name: string;
    location: {
      longitude: number;
      latitude: number;
      altitude: number;
    };
    station_type: string;
    status: 'active' | 'inactive' | 'maintenance';
    last_update: string;
  };
}
