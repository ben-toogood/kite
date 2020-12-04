

/* tslint:disable */
// This file was automatically generated and should not be edited.

// ====================================================
// GraphQL mutation operation: refreshLogin
// ====================================================

export interface refreshLogin_refreshTokens {
  accessToken: string;
  refreshToken: string;
}

export interface refreshLogin {
  refreshTokens: refreshLogin_refreshTokens | null;
}

export interface refreshLoginVariables {
  refreshToken: string;
}


/* tslint:disable */
// This file was automatically generated and should not be edited.

// ====================================================
// GraphQL query operation: getPosts
// ====================================================

export interface getPosts_getPosts_author {
  firstName: string;
  lastName: string;
}

export interface getPosts_getPosts {
  id: string;
  imageURL: string;
  description: string;
  author: getPosts_getPosts_author;
}

export interface getPosts {
  getPosts: (getPosts_getPosts | null)[] | null;
}

export interface getPostsVariables {
  createdBefore?: any | null;
}


/* tslint:disable */
// This file was automatically generated and should not be edited.

// ====================================================
// GraphQL mutation operation: requestLogin
// ====================================================

export interface requestLogin {
  login: boolean | null;
}

export interface requestLoginVariables {
  email: string;
}


/* tslint:disable */
// This file was automatically generated and should not be edited.

// ====================================================
// GraphQL mutation operation: createPost
// ====================================================

export interface createPost_createPost {
  imageURL: string;
}

export interface createPost {
  createPost: createPost_createPost;
}

export interface createPostVariables {
  description: string;
  imageBase64: string;
}

/* tslint:disable */
// This file was automatically generated and should not be edited.

//==============================================================
// START Enums and Input Objects
//==============================================================

//==============================================================
// END Enums and Input Objects
//==============================================================