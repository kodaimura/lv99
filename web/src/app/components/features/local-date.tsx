'use client';

import React, { useEffect, useState } from 'react';

type Props = {
  isoString: string | null | undefined;
};

const LocalDate: React.FC<Props> = ({ isoString }) => {
  const [formatted, setFormatted] = useState<string>('');

  useEffect(() => {
    if (!isoString) {
      setFormatted('');
      return;
    }

    const date = new Date(isoString);

    const year = date.getFullYear();
    const month = String(date.getMonth() + 1).padStart(2, '0');
    const day = String(date.getDate()).padStart(2, '0');
    const hours = String(date.getHours()).padStart(2, '0');
    const minutes = String(date.getMinutes()).padStart(2, '0');

    setFormatted(`${year}-${month}-${day} ${hours}:${minutes}`);
  }, [isoString]);

  if (!isoString) return null;

  return <>{formatted || ''}</>;
};

export default LocalDate;
