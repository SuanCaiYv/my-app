import {createStore, Store} from "vuex";

const store: Store<any> = createStore({
    state() {
        return {
            draftArticleId: String,
        }
    },
    mutations: {
        updatedDraftArticleId(state, draftArticleId) {
            state.draftArticleId = draftArticleId
        }
    },
    getters: {
        draftArticleId: (state) => {
            return state.draftArticleId
        }
    }
})

export default store