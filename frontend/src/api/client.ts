import axios from 'axios';
import {UploadImageRequest, UploadImageResponse} from "./types";

const ax = axios.create({
  baseURL: '/api'
})

export const client = {
  async uploadImage(payload: UploadImageRequest): Promise<UploadImageResponse> {
    let data = new FormData()
    data.append('file', payload.file)

    let config = {
      headers: {
        'Content-Type': payload.contentType
      },
    }

    const resp = await ax.post<UploadImageResponse>('/images', data, config)
    return resp.data
  }
}