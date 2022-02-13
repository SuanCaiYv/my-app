import {createStore, Store} from "vuex";

const store: Store<any> = createStore({
    state() {
        return {
            authed: Boolean,
            accessToken: String,
            refreshToken: String,
            userRole: String,
            logoImg: Image,
            draftArticleId: String,
        }
    },
    mutations: {
        updateAuthed(state, authed) {
            state.authed = authed
        },
        updateAccessToken(state, accessToken) {
            state.accessToken = accessToken
        },
        updateRefreshToken(state, refreshToken) {
            state.refreshToken = refreshToken
        },
        updateUserRole(state, userRole) {
            state.userRole = userRole
        },
        updatedDraftArticleId(state, draftArticleId) {
            state.draftArticleId = draftArticleId
        }
    },
    getters: {
        authed: (state) => {
            return state.authed
        },
        accessToken: (state) => {
            return state.accessToken
        },
        refreshToken: (state) => {
            return state.refreshToken
        },
        userRole: (state) => {
            return state.userRole
        },
        logoImg: (state) => {
            return state.logoImg
        },
        draftArticleId: (state) => {
            return state.draftArticleId
        }
    }
})

export default store