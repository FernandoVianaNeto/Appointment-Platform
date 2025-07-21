import { ScrollableList } from './styles';

type Props = {
  children?: React.ReactNode,
  noContent?: boolean,
};

function DashboardList({ noContent, children }: Props) {
  return (
    <ScrollableList noContent>
        {children}
    </ScrollableList>
  );
}

export default DashboardList;