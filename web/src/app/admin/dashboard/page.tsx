export const dynamic = 'force-dynamic';

import React from 'react';
import { api } from '@/app/lib/api/api.server';

const DashboardPage: React.FC = async () => {
  const data: any = await api.get('admins/me');
  return (
    <div>
      ようこそ {data.admin_name} さん
    </div>
  );
};

export default DashboardPage;
