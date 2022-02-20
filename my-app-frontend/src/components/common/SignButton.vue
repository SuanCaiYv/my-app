<template>
    <div class="sign">
        <button class="button" @click="signButton">{{ state }}</button>
    </div>
</template>

<script setup lang="ts">
import {ref} from "vue"
import {useRouter} from "vue-router";
import storage from "../../util/storage"
import {parseBoolean} from "../../util/base";
import {Constant} from "../../common/systemconstant";

const name = ref<String>("SignButton")

const state = ref<String>("Sign")

// todo
if (parseBoolean(storage.get(Constant.AUTHENTICATED))) {
    state.value = "登出"
} else {
    state.value = "登录"
}

const router = useRouter()
const signButton = function () {
    if (parseBoolean(storage.get(Constant.AUTHENTICATED))) {
        storage.set(Constant.AUTHENTICATED, "false")
        storage.set(Constant.ACCESS_TOKEN, "")
        storage.set(Constant.REFRESH_TOKEN, "")
        state.value = "登录"
    } else {
        router.push("/sign")
    }
}
</script>

<style scoped>

.sign {
    width: 100%;
    height: 100%;
    grid-area: sign-button;
    text-align: left;
}

.button {
    width: auto;
    height: 40px;
    margin-left: 0;
    margin-top: 10px;
    margin-bottom: 10px;
    padding: 0 20px;
    border: none;
    box-sizing: border-box;
    border-radius: 18px;
    font-size: 1.0rem;
    font-weight: bolder;
    background-color: rgba(0,0,0,0.1);
}

.button:hover {
    background-color: rgba(0,0,0,0.15);
}

.button:active {
    background-color: rgba(0,0,0,0.2);
}
</style>