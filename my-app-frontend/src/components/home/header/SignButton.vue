<template>
    <div class="sign-button">
        <button class="sign-button-button" @click="signButton">
            <span style="font-size: 1.2rem">{{ state }}</span>
        </button>
    </div>
</template>

<script setup lang="ts">
import {ref} from "vue"
import {useRouter} from "vue-router";
import {useStore} from "vuex";

const name = ref<String>("SignButton")

const state = ref<String>("Sign")
const store = useStore()

// todo
if (store.getters.authed) {
    state.value = "Logout"
} else {
    state.value = "Sign"
}

const router = useRouter()
const signButton = function () {
    if (store.getters.authed) {
        store.commit("updateAuthed", false)
        store.commit("updateAccessToken", "")
        store.commit("updateRefreshToken", "")
        state.value = "Logout"
    } else {
        router.push("/sign")
    }
}
</script>

<style scoped>

.sign-button {
    width: 100%;
    height: 100%;
    grid-area: sign-button;
}

.sign-button-button {
    width: 120px;
    height: 100%;
    margin-left: 40px;
    margin-right: 40px;
    padding: 0;
    border: 1px solid silver;
    box-sizing: border-box;
    border-radius: 8px;
    background-color: lightskyblue;
}

.sign-button-button:hover {
    background-color: deepskyblue;
}

.sign-button-button:active {
    background-color: dodgerblue;
}
</style>