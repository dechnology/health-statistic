export interface HKData {
  id: string;
  type: string;
  value: string;
  startTimestamp: string;
  endTimestamp: string;
  timezoneID: string;
}

export interface HealthKit {
  id: string;
  startTime: Date;
  endTime: Date;
  data: HKData[];
}
