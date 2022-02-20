import {Configuration, KuefaApiFactory} from "./generated";

const cfg = new Configuration({
  basePath: '/api'
});

export const client = KuefaApiFactory(cfg)

// export const client = {
//   async uploadImage(payload: UploadImageRequest): Promise<UploadImageResponse> {
//     let data = new FormData()
//     data.append('file', payload.file)
//
//     let config = {
//       headers: {
//         'Content-Type': payload.contentType
//       },
//     }
//
//     const resp = await ax.post<UploadImageResponse>('/images', data, config)
//     return resp.data
//   }
// }