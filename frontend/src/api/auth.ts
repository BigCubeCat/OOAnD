import axios from "axios";
import {envApiAddress} from "../utils/env.ts";
import {TAuthDTO} from "../dto/auth_dto.ts";

const url = envApiAddress();

export const apiLogin = async (dto: TAuthDTO) => {
    console.log(url + "/api/auth/login");
    await axios.post(url + "/api/auth/login", {
            id: dto.id || 0,
            username: dto.username,
            password: dto.password || '',
            photo_url: dto.photo_url || '',
            first_name: dto.first_name || '',
            last_name: dto.last_name || '',
        },
        {
            headers: {
                'Content-Type': 'multipart/form-data'
            }
        }
    )
};