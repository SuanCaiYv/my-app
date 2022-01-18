<template>
    <div class="userInfo">
        bbb
    </div>
</template>

<script setup lang="ts">
import {ref} from "vue"

const name = ref<String>("UserInfo")
</script>

<style scoped>
.userInfo {
    width: 100%;
    height: 100%;
    grid-area: userInfo;
}
</style>