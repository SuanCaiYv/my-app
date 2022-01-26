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
import storage from "../../../util/storage"
import {parseBoolean} from "../../../util/base";

const name = ref<String>("SignButton")

const state = ref<String>("Sign")

// todo
if (parseBoolean(storage.get("authed"))) {
    state.value = "Logout"
} else {
    state.value = "Sign"
}

const router = useRouter()
const signButton = function () {
    if (parseBoolean(storage.get("authed"))) {
        storage.set("authed", "false")
        storage.set("accessToken", "")
        storage.set("refreshToken", "")
        state.value = "Sign"
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
    width: 150px;
    height: 100%;
    margin-left: 25px;
    margin-right: 25px;
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