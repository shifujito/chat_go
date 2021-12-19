import { atom, selector } from "recoil";
import { loginInfo } from "./types";
import { recoilPersist } from "recoil-persist";

type Sample = {
  title: string;
};

const initSample: Sample = { title: "hello" };

const { persistAtom } = recoilPersist();

export const sampleState = atom<Sample>({
  key: "sampleState",
  default: initSample,
});

export const titleSelector = selector({
  key: "getState",
  get: ({ get }) => {
    return get(sampleState).title;
  },
});

export const singInUserState = atom<loginInfo>({
  key: "singInUserState",
  default: { id: 0, name: "ユーザーネーム", isLogined: false },
  effects_UNSTABLE: [persistAtom],
});
