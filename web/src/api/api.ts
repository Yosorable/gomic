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
  id: number;
  name: string;
  files: ArchiveFile[];
  cover_url: string;
}

export interface Author {
  id: number;
  name: string;
  cover_url: string;
}

export interface UserResponse {
  id: number;
  user_name: string;
  is_admin: boolean;
  jwt_token?: string;
}

enum APIName {
  SCANNER_START = "/api/scanner/start",
  SCANNER_STATUS = "/api/scanner/status",

  ARCHIVE_AUTHORS = "/api/archive/authors",
  ARCHIVE_BY_AUTHOR_NAME = "/api/archive/author/:name",
  ARCHIVE_ALL = "/api/archive/all",
  ARCHIVE_FILE_BY_ID = "/api/archive/files/:id",

  AUTH_LOGIN = "/auth/login",
}

function FetchBackend(
  input: RequestInfo | URL,
  init?: RequestInit
): Promise<any> {
  return fetch(input, {
    method: "POST",
    ...init,
    headers: {
      Authorization: localStorage.getItem("jwt") ?? "",
    },
  })
    .then((res) => res.json())
    .then((res) => {
      if (res.code === 7000 && res.msg == "Unauthorized") {
        window.location.href = "/#/login";
        return;
      }
      return res;
    });
}

export default {
  startScanner(): Promise<ResponseBase> {
    return FetchBackend(APIName.SCANNER_START);
  },
  scannerStatus(): Promise<
    ResponseWithData<{ status: boolean; records: string[] }>
  > {
    return FetchBackend(APIName.SCANNER_STATUS);
  },
  getAuthors(): Promise<ResponseWithData<Author[]>> {
    return FetchBackend(APIName.ARCHIVE_AUTHORS);
  },
  getArchiveByAuthorName(
    name: string,
    page: number = 1,
    limit: number = 50
  ): Promise<ResponseWithData<Archive[]>> {
    return FetchBackend(
      APIName.ARCHIVE_BY_AUTHOR_NAME.replace(":name", name) +
        `?page=${page}&limit=${limit}`
    );
  },
  getAllArchives(
    page: number = 1,
    limit: number = 50
  ): Promise<ResponseWithData<Archive[]>> {
    return FetchBackend(APIName.ARCHIVE_ALL + `?page=${page}&limit=${limit}`);
  },
  getArchiveByID(id: string): Promise<ResponseWithData<Archive>> {
    return FetchBackend(APIName.ARCHIVE_FILE_BY_ID.replace(":id", id));
  },
  authLogin(
    username: string,
    pwd: string
  ): Promise<ResponseWithData<UserResponse>> {
    return FetchBackend(APIName.AUTH_LOGIN, {
      body: JSON.stringify({
        user_name: username,
        password: pwd,
      }),
    });
  },
};
