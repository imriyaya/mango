import { defineStore } from 'pinia'

export const useViewingPage = defineStore('viewingPage', {
    state: () => {
        return { is: false }
    },
    actions: {
        switch(is: boolean) {
            this.is = is
        },
    }
})
