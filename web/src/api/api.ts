export interface ResponseBase {
  code: number;
  msg?: string;
}

export interface ResponseWithData<T> extends ResponseBase {
  data: T;
}

export interface ArchiveFile {
  id: string;
  name: string;
  path: string;
  url: string;
}

export interface Archive {
  name: string;
  files: ArchiveFile[];
  cover_url: string;
}

enum APIName {
  SCANNER_START = "/api/scanner/start",
  SCANNER_STATUS = "/api/scanner/status",
  ARCHIVE_AUTHORS = "/api/archive/authors",
  ARCHIVE_BY_AUTHOR_NAME = "/api/archive/author/:name",
  ARCHIVE_ALL = "/api/archive/all",
  ARCHIVE_BY_NAME = "/api/archive/name/:name",
}

export default {
  startScanner(): Promise<ResponseBase> {
    return fetch(APIName.SCANNER_START, { method: "POST" }).then((res) =>
      res.json()
    );
  },
  scannerStatus(): Promise<
    ResponseWithData<{ status: boolean; records: string[] }>
  > {
    return fetch(APIName.SCANNER_STATUS).then((res) => res.json());
  },
  getAuthors(): Promise<ResponseWithData<string[]>> {
    return fetch(APIName.ARCHIVE_AUTHORS).then((res) => res.json());
  },
  getArchiveByAuthorName(name: string): Promise<ResponseWithData<Archive[]>> {
    return fetch(APIName.ARCHIVE_BY_AUTHOR_NAME.replace(":name", name)).then(
      (res) => res.json()
    );
  },
  getAllArchives(): Promise<ResponseWithData<Archive[]>> {
    return fetch(APIName.ARCHIVE_ALL).then((res) => res.json());
  },
  getArchiveByName(name: string): Promise<ResponseWithData<Archive>> {
    return fetch(APIName.ARCHIVE_BY_NAME.replace(":name", name)).then((res) =>
      res.json()
    );
  },
};
