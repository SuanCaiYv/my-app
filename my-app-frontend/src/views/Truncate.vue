<template>
    <div class="truncate">
    </div>
</template>

<script setup lang="ts">
import {ref} from "vue"
import alertFunc from "../util/alert";
import {useRouter} from "vue-router";
import storage from "../util/storage";

const name = ref<string>("Truncate")
const router = useRouter()

storage.removeAll()

alertFunc("清空数据完成", function () {
    router.push("/home")
})
</script>

<style scoped>
.truncate {
}
</style>