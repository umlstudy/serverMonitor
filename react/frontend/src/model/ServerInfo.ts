import { AbstractStatus } from './AbstractStatus';
import { ProcessStatus } from './ProcessStatus';
import { ResourceStatus } from './ResourceStatus';

// tslint:disable-next-line:interface-name
export interface ServerInfo extends AbstractStatus {
    sortOrder:number;
    resourceStatusesModified:boolean;
    processStatusesModified:boolean;
    resourceStatuses: ResourceStatus[];
    processStatuses: ProcessStatus[];
    isRunning:boolean;
}

export type ServerInfoMap = Record<string, ServerInfo>;
