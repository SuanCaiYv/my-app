import {createStore, Store} from "vuex";

const store: Store<any> = createStore({
    state() {
        return {
            draftArticleId: String,
            operation: String,
        }
    },
    mutations: {
        updateDraftArticleId(state, draftArticleId) {
            state.draftArticleId = draftArticleId
        },
        updateOperation(state, operation) {
            state.operation = operation
        },
    },
    getters: {
        draftArticleId: (state) => {
            return state.draftArticleId
        },
        operation: (state) => {
            return state.operation
        },
    }
})

export default store