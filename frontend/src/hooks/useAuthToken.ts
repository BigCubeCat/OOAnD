import {useAppSelector} from '../store/hooks';

import {selectToken} from "../store/slices/userSlice.ts";

export default function useAuthToken() {
    const token = useAppSelector(selectToken);
    return {
        token,
        authorized: token != "",
    }
}