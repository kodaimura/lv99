import { Account } from "@/types/models";
import { createContext, useContext, useState, ReactNode } from "react";

type AuthContextType = {
  account: Account | null;
  setAccount: (account: Account | null) => void;
};

const AuthContext = createContext<AuthContextType | undefined>(undefined);

type Props = {
  children: ReactNode;
  initAccount: Account | null;
};

export const AuthProvider: React.FC<Props> = ({ children, initAccount }) => {
  const [account, setAccount] = useState<Account | null>(initAccount);

  return (
    <AuthContext.Provider value={{ account, setAccount }}>
      {children}
    </AuthContext.Provider>
  );
};

export const useAuth = (): AuthContextType => {
  const context = useContext(AuthContext);
  if (context === undefined) {
    throw new Error("useAuth must be used within an AuthProvider");
  }
  return context;
};
