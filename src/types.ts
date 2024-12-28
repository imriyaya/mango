export class Manga {
    constructor(
        public _id: string,
        public title: string,
        public chapters: Chapter[]
    ) {
    }
}

export class Chapter {
    constructor(
        public id: number,
        public title: string,
        public page: number
    ) {
    }
}

export class MangaListResponse {
    constructor(
        public count: number,
        public result: Manga[]
    ) {
    }
}