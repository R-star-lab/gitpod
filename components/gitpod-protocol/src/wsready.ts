// generated using github.com/32leaves/bel on 2021-01-11 14:48:40.747399219 +0000 UTC m=+0.007952821
// DO NOT MODIFY

export enum WorkspaceInitSource {
    WorkspaceInitFromBackup = "from-backup",
    WorkspaceInitFromPrebuild = "from-prebuild",
    WorkspaceInitFromOther = "from-other",
}
export interface WorkspaceReadyMessage {
    source: WorkspaceInitSource
}
