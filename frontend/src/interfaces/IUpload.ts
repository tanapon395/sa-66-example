export interface ImageUpload {
    uid: string
    lastModified: number
    lastModifiedDate: string
    name: string
    size: number
    type: string
    percent: number
    originFileObj: OriginFileObj
    error: Error
    response: string
    status: string
    thumbUrl: string
  }
  
  export interface OriginFileObj {
    uid: string
  }
  
  export interface Error {
    status: number
    method: string
    url: string
  }