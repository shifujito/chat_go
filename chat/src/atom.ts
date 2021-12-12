import { atom, selector } from "recoil";
import { loginInfo } from './types'

type Sample = {
    title: string;
}

const initSample: Sample = {title: "hello"}

export const sampleState = atom<Sample>({
    key: 'sampleState',
    default: initSample
})

export const titleSelector = selector({
    key: 'getState',
    get: ({get}) => {
        return  get(sampleState).title
    }
})

export const singInUserState = atom<loginInfo>({
    key: 'singInUserState',
    default: {name: '', isLogined: false}
})
