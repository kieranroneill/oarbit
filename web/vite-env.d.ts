/// <reference types="vite/client" />

declare namespace NodeJS {
  interface ProcessEnv {
    readonly API_BASE_URL?: string;
  }
}

interface ImportMetaEnv {
  readonly VITE_DESCRIPTION: string;
  readonly VITE_LOG_LEVEL: string;
  readonly VITE_TITLE: string;
}

interface ImportMeta {
  readonly env: ImportMetaEnv;
}
