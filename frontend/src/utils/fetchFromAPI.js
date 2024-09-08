import axios from 'axios';

const metadata_url = 'http://localhost:8002';
const videofile_url = 'http://localhost:8001';

const headers = {
    'Content-Type': 'application/json',
    'Accept': 'application/json',
  };
  
export const getAllMetadata = async () => {
    const { data } = await axios.get(`${metadata_url}/videos`, { headers });
    return data;
} 

export const getVideoMetadataByID = async (id) => {
    const { data } = await axios.get(`${metadata_url}/videos/${id}`, { headers });
    return data;
}

export const getVideoMetadataByTag = async (tag) => {
    const { data } = await axios.get(`${metadata_url}/video/tag/${tag}`, { headers });
    return data;
}

export const getVideo = async (filename) => {
    const { data } = await axios.get(`${videofile_url}/${filename}`, { headers });
    return data;
}
