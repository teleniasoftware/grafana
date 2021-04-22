import React, { FC } from 'react';
import { LinkButton, VerticalGroup } from '@grafana/ui';
import { css } from '@emotion/css';
import { getConfig } from 'app/core/config';

export const UserSignup: FC<{}> = () => {
  const href = getConfig().verifyEmailEnabled ? `${getConfig().appSubUrl}/verify` : `${getConfig().appSubUrl}/signup`;
  return (
    <VerticalGroup
      className={css`
        margin-top: 16px;
      `}
    >
      <span>New to Grafana?</span>
      <LinkButton
        className={css`
          width: 100%;
          justify-content: center;
        `}
        href={href}
        variant="secondary"
      >
        Sign Up
      </LinkButton>
    </VerticalGroup>
  );
};
