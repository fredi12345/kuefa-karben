export type UploadImageRequest = {
  isTitle: boolean,
  contentType: string,
  file: any,
}

export type UploadImageResponse = {
  id: string,
  imageURL: string,
  thumbnailURL: string,
}
