export type TUser = {
    id: number;
    tgId: number;
    handle: string;
    avatar: string;
};

export const defaultUser = {
    id: 0,
    tgId: 0,
    handle: '',
    avatar: '',
};