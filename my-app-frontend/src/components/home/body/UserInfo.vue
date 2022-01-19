<template>
    <div class="userInfo">
        <div class="img">
            <img style="width: 100%; height: 100%; object-fit: fill" src="../../../assets/IMG_0751.JPG" alt="no-avatar">
        </div>
        <div class="nickname">
            {{nickname}}
        </div>
        <div class="email">
            {{email}}
        </div>
        <div class="signature">
            {{signature}}
        </div>
    </div>
</template>

<script setup lang="ts">
import {ref} from "vue"

const name = ref<String>("UserInfo")

const props = defineProps({
    nickname: String,
    email: String,
    signature: String
})
</script>

<style scoped>
.userInfo {
    width: 100%;
    height: 100%;
    grid-area: userInfo;
    border: 1px solid silver;
    box-sizing: border-box;
}

.img {
    width: 150px;
    height: 150px;
    margin: 25px 25px 0;
    padding: 0;
    border: 1px solid white;
    box-sizing: border-box;
    border-radius: 6px;
}

.nickname {
    width: 150px;
    height: 20px;
    margin: 25px 25px 0;
    padding: 0;
    border: 1px solid black;
    box-sizing: border-box;
    overflow-x: scroll;
    border-radius: 6px;
    font-size: 1rem;
    line-height: 20px;
}

.email {
    width: 180px;
    height: 20px;
    margin: 25px 10px 0;
    padding: 0;
    border: 1px solid black;
    box-sizing: border-box;
    border-radius: 6px;
    overflow-x: scroll;
    font-size: 1rem;
    line-height: 20px;
}

.signature {
    width: 180px;
    height: auto;
    margin: 25px 10px 0;
    padding: 0;
    border: 1px solid black;
    box-sizing: border-box;
    border-radius: 4px;
    font-size: 1rem;
}
</style>