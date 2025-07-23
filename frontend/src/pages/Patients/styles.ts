import styled from 'styled-components';

export const Container = styled.div`
  display: flex;
  margin-bottom: 0px;
`;

export const H1 = styled.h1`
  font-family: ${({ theme }) => theme.font.family};
  color: ${({ theme }) => theme.colors.primary};
`;

export const Wrapper = styled.div`
  display: flex;
  align-items: center;
  justify-content: space-between;
  margin-top: 30px;
`;

export const Div = styled.div`
  display: flex;
  gap: 20px;
`;

export const ListOptionsWrapper = styled.div<{
  deleteSelection?: boolean;
}>`
  display: flex;
  justify-content: space-between;
  align-items: center;

  button {
    visibility: ${({ deleteSelection }) => (deleteSelection ? 'visible' : 'hidden')};
    display: flex;
    align-items: center;
    justify-content: center;
    border: none;
    background: transparent;
    cursor: pointer;
    font-size: 20px;
    margin: 0;
    padding: 0;
    color: red;
  }
`;

export const DashboardWrapper = styled.div`
  display: flex;
  flex-direction: column;
  margin-top: 50px;
  gap: 20px;

  p {
    font-size: 16px;
    font-weight: bold;
    color: ${({ theme }) => theme.colors.primary}
  }

  span {
    display: flex;
    align-items: center;
    gap: 5px;
  }
`;