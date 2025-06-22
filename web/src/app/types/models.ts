export type Account = {
  id: number,
  name: string,
  role: number,
  created_at: string,
  updated_at: string,
}

export type AccountProfile = {
  account_id: number,
  display_name: string,
  bio: string,
  avatar_url: string,
  created_at: string,
  updated_at: string,
}

export type AccountWithProfile = {
  id: number,
  name: string,
  role: number,
  display_name: string,
  bio: string,
  avatar_url: string,
  created_at: string,
  updated_at: string,
  deleted_at: string | null,
};

export type Question = {
  id: number;
  title: string;
  content: string;
  answer: string;
  level: number;
  updated_at: string;
  deleted_at: string | null;
};

export type Answer = {
  id: number;
  question_id: number;
  account_id: number;
  code_def: string;
  code_call: string;
  call_output: string;
  call_error: string;
  is_correct: boolean | null;
  correct_at: string | null;
  updated_at: string;
};

export type Comment = {
  id: number;
  answer_id: number;
  account_id: number;
  content: string;
  created_at: string;
  updated_at: string;
};

export type Chat = {
  id: number;
  from_id: number;
  to_id: number;
  message: string;
  is_read: boolean;
  created_at: string;
  updated_at: string;
};