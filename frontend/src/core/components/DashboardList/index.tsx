import { ScrollableList } from './styles';

type Props = {
  children?: React.ReactNode,
};

function DashboardList({ children }: Props) {
  return (
    <ScrollableList>
        {children}
    </ScrollableList>
  );
}

export default DashboardList;