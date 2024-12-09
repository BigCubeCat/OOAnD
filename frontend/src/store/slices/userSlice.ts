import {createSlice, PayloadAction} from '@reduxjs/toolkit';
import {RootState} from '../store';
import {TUser} from '../types';

export interface IUserState {
    user: TUser;
}

// TODO add fetch
const initialState: IUserState = {
    user: {
        name: 'guest',
        age: 40,
        experience: 5,
        count: 1,
        income: 1000000,
        consumption: 10000,
        phone: '+7',
        region: '22',
        provision: 'none',
        inn: '0000000000',
    },
};

export const userSlice = createSlice({
    name: 'user',
    initialState,
    reducers: {
        setUser: (state, action: PayloadAction<TUser>) => {
            state.user = action.payload;
        },
        setName: (state, action: PayloadAction<string>) => {
            state.user.name = action.payload;
        },
        setAge: (state, action: PayloadAction<number>) => {
            state.user.age = action.payload;
        },
        setExp: (state, action: PayloadAction<number>) => {
            state.user.experience = action.payload;
        },
        setCount: (state, action: PayloadAction<number>) => {
            state.user.count = action.payload;
        },
        setIncome: (state, action: PayloadAction<number>) => {
            state.user.income = action.payload;
        },
        setConsumption: (state, action: PayloadAction<number>) => {
            state.user.consumption = action.payload;
        },
        setPhone: (state, action: PayloadAction<string>) => {
            state.user.phone = action.payload;
        },
        setInn: (state, action: PayloadAction<string>) => {
            state.user.inn = action.payload;
        },
        setRegion: (state, action: PayloadAction<string>) => {
            state.user.region = action.payload;
        },
        setProvision: (state, action: PayloadAction<string>) => {
            state.user.provision = action.payload;
        },
    },
});

export const {
    setUser,
    setProvision,
    setInn,
    setRegion,
    setExp,
    setPhone,
    setConsumption,
    setName,
    setIncome,
    setAge,
    setCount,
} = userSlice.actions;

export const selectUser = (state: RootState) => state.user;
export default userSlice.reducer;