import superagent from "superagent";
import {URL_API, URL_ENDPOINT} from "../constants/URLConstants";

export const get = async () => {
    try {
        const res = await superagent.get(`${URL_API}${URL_ENDPOINT}`);
        return JSON.parse(res.text);
    } catch(err) {
        console.log({err})
    }
}

export const post = async (customer) => {
    console.log({customer})
    try {
        const res = await superagent.post(`${URL_API}${URL_ENDPOINT}`).send(customer);
        return JSON.parse(res.text);
    } catch(err) {
        throw err;
    }
}