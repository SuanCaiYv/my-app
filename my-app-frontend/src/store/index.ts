import {createStore, Store} from "vuex";

const store: Store<any> = createStore({
    state() {
        return {
            accessToken: String,
            refreshToken: String,
            userRole: String
        }
    },
    mutations: {
        updateAccessToken(state, accessToken) {
            state.accessToken = accessToken
        },
        updateRefreshToken(state, refreshToken) {
            state.refreshToken = refreshToken
        },
        updateUserRole(state, userRole) {
            state.userRole = userRole
        }
    },
    getters: {
        accessToken: (state) => {
            return state.accessToken
        },
        refreshToken: (state) => {
            return state.refreshToken
        },
        userRole: (state) => {
            return state.userRole
        }
    }
})

export default store