export interface ReqPage {
    page: number;
    pageSize: number;
}

export interface CommonModel {
    id: number;
    CreatedAt?: string;
    UpdatedAt?: string;
}

export interface ResultData<T> {
    code: number;
    message: string;
    data: T;
}

export interface ResPage<T> {
    items: T[];
    total: number;
}
