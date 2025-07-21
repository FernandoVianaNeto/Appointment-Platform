import InfiniteScroll from 'react-infinite-scroll-component';
import { ScrollableList } from './styles';
import LoadingSpinner from '../Loading';

type Props = {
  children?: React.ReactNode,
  noContent?: boolean,
  fetchMoreData: () => void;
  hasMore: boolean;
};

function DashboardList({ noContent, children, hasMore, fetchMoreData }: Props) {
  return (
    <ScrollableList id="dashboardScrollContainer" noContent={noContent}>
      <InfiniteScroll
        dataLength={10}
        next={fetchMoreData}
        hasMore={hasMore}
        loader={<LoadingSpinner />}
        scrollableTarget="dashboardScrollContainer"
        style={{ overflow: 'visible' }}
      >
        {children}
      </InfiniteScroll>
    </ScrollableList>
  );
}

export default DashboardList;