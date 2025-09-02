import {createSlice, PayloadAction} from '@reduxjs/toolkit';
import {RootState} from '../store';
import {defaultUser, TUser} from '../types';

export interface IUserState {
    user: TUser;
    token: string;
}

const initialState: IUserState = {
    user: defaultUser,
    token: ''
};

export const userSlice = createSlice({
    name: 'user',
    initialState,
    reducers: {
        setUser: (state, action: PayloadAction<TUser>) => {
            state.user = action.payload;
        },
    },
});

export const {
    setUser,
} = userSlice.actions;

export const selectUser = (state: RootState) => state.user.user;
export const selectToken = (state: RootState) => state.user.token;
export default userSlice.reducer;