import ProductSimple from '../components/Cards'
import { SimpleGrid, Grid, GridItem, Text } from '@chakra-ui/react'
export default function Products() {
    return (
        <>
            <Grid
                templateAreas={`"new"
                  "promotion"
                  "all"`}
                gridTemplateRows={'50px 1fr 30px'}
                gridTemplateColumns={'1fr'}
                gap='1'
                fontWeight='bold'
            ></Grid>
            <GridItem pl='2' area={'new'}>
                <Text>Sản phẩm Mới</Text>
                <SimpleGrid columns={[2, null, 4]} spacing='40px'>
                    <ProductSimple />
                    <ProductSimple />
                    <ProductSimple />
                </SimpleGrid>
            </GridItem>
            <GridItem pl='2' bg='orange.300' area={'promotion'}>
                <Text>Khuyến mãi</Text>
                <SimpleGrid columns={[2, null, 4]} spacing='40px'>
                    <ProductSimple />
                    <ProductSimple />
                    <ProductSimple />
                </SimpleGrid>
            </GridItem>
            <GridItem pl='2' area={'all'}>
                <Text>Tất cả sản phẩm</Text>
                <SimpleGrid columns={[2, null, 4]} spacing='40px'>
                    <ProductSimple />
                    <ProductSimple />
                    <ProductSimple />
                    <ProductSimple />
                    <ProductSimple />
                    <ProductSimple />
                    <ProductSimple />
                    <ProductSimple />
                    <ProductSimple />
                    <ProductSimple />
                    <ProductSimple />
                    <ProductSimple />
                    <ProductSimple />
                    <ProductSimple />
                    <ProductSimple />
                    <ProductSimple />
                    <ProductSimple />


                </SimpleGrid>
            </GridItem>
        </>
    )
}