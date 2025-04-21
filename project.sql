create table Blog.Articles
(
    Id           bigint auto_increment
        primary key,
    Title        longtext                                  null,
    Description  longtext                                  null,
    Body         longtext                                  null,
    Slug         longtext                                  not null,
    SiteId       bigint                                    not null,
    VisitedCount int                                       not null,
    ReviewCount  int                                       not null,
    Rate         int                                       not null,
    Badges       longtext                                  null,
    SeoTags      longtext                                  null,
    UserId       bigint                                    not null,
    CreatedAt    datetime(6)                               not null,
    UpdatedAt    datetime(6)                               not null,
    Version      timestamp(6) default current_timestamp(6) not null on update current_timestamp(6),
    IsDeleted    tinyint(1)                                not null,
    DeletedAt    datetime(6)                               null
);

create table Blog.ArticleMedia
(
    Id        bigint auto_increment
        primary key,
    ArticleId bigint not null,
    MediaId   bigint not null,
    constraint FK_ArticleMedia_Articles_ArticleId
        foreign key (ArticleId) references Blog.Articles (Id)
            on delete cascade
);

create index IX_ArticleMedia_ArticleId
    on Blog.ArticleMedia (ArticleId);

create table `Order`.Baskets
(
    Id                           bigint auto_increment
        primary key,
    SiteId                       bigint                                    not null,
    TotalRawPrice                bigint                                    not null,
    TotalCouponDiscount          bigint                                    not null,
    TotalPriceWithCouponDiscount bigint                                    not null,
    DiscountId                   bigint                                    null,
    CustomerId                   bigint                                    not null,
    CreatedAt                    datetime(6)                               not null,
    UpdatedAt                    datetime(6)                               not null,
    Version                      timestamp(6) default current_timestamp(6) not null on update current_timestamp(6),
    IsDeleted                    tinyint(1)                                not null,
    DeletedAt                    datetime(6)                               null
);

create table `Order`.BasketItems
(
    Id                           bigint auto_increment
        primary key,
    Quantity                     int                                       not null,
    RawPrice                     bigint                                    not null,
    FinalRawPrice                bigint                                    not null,
    FinalPriceWithCouponDiscount bigint                                    not null,
    JustCouponPrice              bigint                                    not null,
    JustDiscountPrice            bigint                                    not null,
    BasketId                     bigint                                    not null,
    ProductId                    bigint                                    not null,
    ProductVariantId             bigint                                    not null,
    CreatedAt                    datetime(6)                               not null,
    UpdatedAt                    datetime(6)                               not null,
    Version                      timestamp(6) default current_timestamp(6) not null on update current_timestamp(6),
    IsDeleted                    tinyint(1)                                not null,
    DeletedAt                    datetime(6)                               null,
    constraint FK_BasketItems_Baskets_BasketId
        foreign key (BasketId) references `Order`.Baskets (Id)
            on delete cascade
);

create index IX_BasketItems_BasketId
    on `Order`.BasketItems (BasketId);

create table Product.Categories
(
    Id               bigint auto_increment
        primary key,
    Name             longtext                                  not null,
    ParentCategoryId bigint                                    null,
    `Order`          int                                       not null,
    Description      longtext                                  null,
    Slug             longtext                                  not null,
    SeoTags          longtext                                  null,
    SiteId           bigint                                    not null,
    CategoryId       bigint                                    null,
    UserId           bigint                                    not null,
    CreatedAt        datetime(6)                               not null,
    UpdatedAt        datetime(6)                               not null,
    Version          timestamp(6) default current_timestamp(6) not null on update current_timestamp(6),
    IsDeleted        tinyint(1)                                not null,
    DeletedAt        datetime(6)                               null,
    constraint FK_Categories_Categories_CategoryId
        foreign key (CategoryId) references Product.Categories (Id)
);

create table Blog.Categories
(
    Id               bigint auto_increment
        primary key,
    Name             longtext                                  not null,
    ParentCategoryId bigint                                    null,
    `Order`          int                                       not null,
    Description      longtext                                  null,
    Slug             longtext                                  not null,
    SeoTags          longtext                                  null,
    SiteId           bigint                                    not null,
    UserId           bigint                                    not null,
    CreatedAt        datetime(6)                               not null,
    UpdatedAt        datetime(6)                               not null,
    Version          timestamp(6) default current_timestamp(6) not null on update current_timestamp(6),
    IsDeleted        tinyint(1)                                not null,
    DeletedAt        datetime(6)                               null,
    constraint FK_Categories_Categories_ParentCategoryId
        foreign key (ParentCategoryId) references Blog.Categories (Id)
);

create table Blog.ArticleCategory
(
    Id         bigint auto_increment
        primary key,
    ArticleId  bigint not null,
    CategoryId bigint not null,
    constraint FK_ArticleCategory_Articles_ArticleId
        foreign key (ArticleId) references Blog.Articles (Id)
            on delete cascade,
    constraint FK_ArticleCategory_Categories_CategoryId
        foreign key (CategoryId) references Blog.Categories (Id)
            on delete cascade
);

create index IX_ArticleCategory_ArticleId
    on Blog.ArticleCategory (ArticleId);

create index IX_ArticleCategory_CategoryId
    on Blog.ArticleCategory (CategoryId);

create index IX_Categories_CategoryId
    on Product.Categories (CategoryId);

create index IX_Categories_ParentCategoryId
    on Blog.Categories (ParentCategoryId);

create table Product.CategoryMedia
(
    Id         bigint auto_increment
        primary key,
    CategoryId bigint not null,
    MediaId    bigint not null,
    constraint FK_CategoryMedia_Categories_CategoryId
        foreign key (CategoryId) references Product.Categories (Id)
            on delete cascade
);

create table Blog.CategoryMedia
(
    Id         bigint auto_increment
        primary key,
    CategoryId bigint not null,
    MediaId    bigint not null,
    constraint FK_CategoryMedia_Categories_CategoryId
        foreign key (CategoryId) references Blog.Categories (Id)
            on delete cascade
);

create index IX_CategoryMedia_CategoryId
    on Product.CategoryMedia (CategoryId);

create index IX_CategoryMedia_CategoryId
    on Blog.CategoryMedia (CategoryId);

create table Ai.Credits
(
    Id         bigint auto_increment
        primary key,
    UserId     bigint                                    not null,
    CustomerId bigint                                    not null,
    CreatedAt  datetime(6)                               not null,
    UpdatedAt  datetime(6)                               not null,
    Version    timestamp(6) default current_timestamp(6) not null on update current_timestamp(6),
    IsDeleted  tinyint(1)                                not null,
    DeletedAt  datetime(6)                               null
);

create table Support.CustomerTickets
(
    Id         bigint auto_increment
        primary key,
    Title      longtext                                  not null,
    Status     longtext                                  not null,
    Category   longtext                                  not null,
    AssignedTo bigint                                    null,
    ClosedBy   bigint                                    null,
    ClosedAt   datetime(6)                               null,
    Priority   longtext                                  not null,
    UserId     bigint                                    not null,
    CustomerId bigint                                    not null,
    CreatedAt  datetime(6)                               not null,
    UpdatedAt  datetime(6)                               not null,
    Version    timestamp(6) default current_timestamp(6) not null on update current_timestamp(6),
    IsDeleted  tinyint(1)                                not null,
    DeletedAt  datetime(6)                               null
);

create table Support.CustomerComments
(
    Id               bigint auto_increment
        primary key,
    CustomerTicketId bigint                                    not null,
    Content          longtext                                  not null,
    RespondentId     bigint                                    not null,
    CreatedAt        datetime(6)                               not null,
    UpdatedAt        datetime(6)                               not null,
    Version          timestamp(6) default current_timestamp(6) not null on update current_timestamp(6),
    IsDeleted        tinyint(1)                                not null,
    DeletedAt        datetime(6)                               null,
    constraint FK_CustomerComments_CustomerTickets_CustomerTicketId
        foreign key (CustomerTicketId) references Support.CustomerTickets (Id)
            on delete cascade
);

create index IX_CustomerComments_CustomerTicketId
    on Support.CustomerComments (CustomerTicketId);

create table Support.CustomerTicketMedia
(
    Id               bigint auto_increment
        primary key,
    CustomerTicketId bigint not null,
    MediaId          bigint not null,
    constraint FK_CustomerTicketMedia_CustomerTickets_CustomerTicketId
        foreign key (CustomerTicketId) references Support.CustomerTickets (Id)
            on delete cascade
);

create index IX_CustomerTicketMedia_CustomerTicketId
    on Support.CustomerTicketMedia (CustomerTicketId);

create table User.Customers
(
    Id                 bigint auto_increment
        primary key,
    SiteId             bigint                                    not null,
    FirstName          longtext                                  null,
    AvatarId           bigint                                    null,
    LastName           longtext                                  null,
    Email              varchar(255)                              not null,
    VerifyEmail        longtext                                  null,
    Password           longtext                                  not null,
    Salt               longtext                                  not null,
    NationalCode       longtext                                  null,
    Phone              longtext                                  null,
    VerifyPhone        longtext                                  null,
    IsActive           longtext                                  not null,
    VerifyCode         int                                       null,
    ExpireVerifyCodeAt datetime(6)                               null,
    CreatedAt          datetime(6)                               not null,
    UpdatedAt          datetime(6)                               not null,
    Version            timestamp(6) default current_timestamp(6) not null on update current_timestamp(6),
    IsDeleted          tinyint(1)                                not null,
    DeletedAt          datetime(6)                               null,
    constraint IX_Customers_Email
        unique (Email)
);

create table Site.DefaultThemes
(
    Id          bigint auto_increment
        primary key,
    Name        longtext                                  not null,
    Description longtext                                  null,
    Demo        longtext                                  null,
    MediaId     bigint                                    not null,
    Pages       longtext                                  not null,
    CreatedAt   datetime(6)                               not null,
    UpdatedAt   datetime(6)                               not null,
    Version     timestamp(6) default current_timestamp(6) not null on update current_timestamp(6),
    IsDeleted   tinyint(1)                                not null,
    DeletedAt   datetime(6)                               null
);

create table Product.Discounts
(
    Id         bigint auto_increment
        primary key,
    Code       longtext                                  not null,
    Quantity   int                                       not null,
    Type       longtext                                  not null,
    Value      bigint                                    not null,
    ExpiryDate datetime(6)                               not null,
    SiteId     bigint                                    not null,
    UserId     bigint                                    not null,
    CreatedAt  datetime(6)                               not null,
    UpdatedAt  datetime(6)                               not null,
    Version    timestamp(6) default current_timestamp(6) not null on update current_timestamp(6),
    IsDeleted  tinyint(1)                                not null,
    DeletedAt  datetime(6)                               null
);

create table Product.CustomerDiscount
(
    Id         bigint auto_increment
        primary key,
    DiscountId bigint not null,
    CustomerId bigint not null,
    constraint FK_CustomerDiscount_Discounts_DiscountId
        foreign key (DiscountId) references Product.Discounts (Id)
            on delete cascade
);

create index IX_CustomerDiscount_DiscountId
    on Product.CustomerDiscount (DiscountId);

create table Drive.FileItems
(
    Id          bigint auto_increment
        primary key,
    Name        longtext                                  not null,
    BucketName  longtext                                  not null,
    ServerKey   longtext                                  not null,
    FilePath    longtext                                  not null,
    IsDirectory tinyint(1)                                not null,
    Size        bigint                                    not null,
    MimeType    longtext                                  not null,
    ParentId    bigint                                    null,
    Permission  longtext                                  not null,
    UserId      bigint                                    not null,
    CreatedAt   datetime(6)                               not null,
    UpdatedAt   datetime(6)                               not null,
    Version     timestamp(6) default current_timestamp(6) not null on update current_timestamp(6),
    IsDeleted   tinyint(1)                                not null,
    DeletedAt   datetime(6)                               null,
    constraint FK_FileItems_FileItems_ParentId
        foreign key (ParentId) references Drive.FileItems (Id)
);

create index IX_FileItems_ParentId
    on Drive.FileItems (ParentId);

create table Payment.Gateways
(
    Id                                   bigint auto_increment
        primary key,
    SiteId                               bigint                                    not null,
    Saman_MerchantId                     longtext                                  null,
    Saman_Password                       longtext                                  null,
    IsActiveSaman                        longtext                                  not null,
    Mellat_TerminalId                    bigint                                    null,
    Mellat_UserName                      longtext                                  null,
    Mellat_UserPassword                  longtext                                  null,
    IsActiveMellat                       longtext                                  not null,
    Parsian_LoginAccount                 longtext                                  null,
    IsActiveParsian                      longtext                                  not null,
    Pasargad_MerchantCode                longtext                                  null,
    Pasargad_TerminalCode                longtext                                  null,
    Pasargad_PrivateKey                  longtext                                  null,
    IsActivePasargad                     longtext                                  not null,
    IranKish_TerminalId                  longtext                                  null,
    IranKish_AcceptorId                  longtext                                  null,
    IranKish_PassPhrase                  longtext                                  null,
    IranKish_PublicKey                   longtext                                  null,
    IsActiveIranKish                     longtext                                  not null,
    Melli_TerminalId                     longtext                                  null,
    Melli_MerchantId                     longtext                                  null,
    Melli_TerminalKey                    longtext                                  null,
    IsActiveMelli                        longtext                                  not null,
    AsanPardakht_MerchantConfigurationId longtext                                  null,
    AsanPardakht_UserName                longtext                                  null,
    AsanPardakht_Password                longtext                                  null,
    AsanPardakht_Key                     longtext                                  null,
    AsanPardakht_IV                      longtext                                  null,
    IsActiveAsanPardakht                 longtext                                  not null,
    Sepehr_TerminalId                    bigint                                    null,
    IsActiveSepehr                       longtext                                  not null,
    ZarinPal_MerchantId                  longtext                                  null,
    ZarinPal_AuthorizationToken          longtext                                  null,
    ZarinPal_IsSandbox                   tinyint(1)                                null,
    IsActiveZarinPal                     longtext                                  not null,
    PayIr_Api                            longtext                                  null,
    PayIr_IsTestAccount                  tinyint(1)                                null,
    IsActivePayIr                        longtext                                  not null,
    IdPay_Api                            longtext                                  null,
    IdPay_IsTestAccount                  tinyint(1)                                null,
    IsActiveIdPay                        longtext                                  not null,
    YekPay_MerchantId                    longtext                                  null,
    IsActiveYekPay                       longtext                                  not null,
    PayPing_AccessToken                  longtext                                  null,
    IsActivePayPing                      longtext                                  not null,
    IsActiveParbadVirtual                longtext                                  not null,
    UserId                               bigint                                    not null,
    CreatedAt                            datetime(6)                               not null,
    UpdatedAt                            datetime(6)                               not null,
    Version                              timestamp(6) default current_timestamp(6) not null on update current_timestamp(6),
    IsDeleted                            tinyint(1)                                not null,
    DeletedAt                            datetime(6)                               null
);

create table Site.HeaderFooters
(
    Id        bigint auto_increment
        primary key,
    SiteId    bigint                                    not null,
    Title     longtext                                  not null,
    IsMain    tinyint(1)                                not null,
    Body      longtext collate utf8mb4_bin              null
        check (json_valid(`Body`)),
    Type      int                                       not null,
    UserId    bigint                                    not null,
    CreatedAt datetime(6)                               not null,
    UpdatedAt datetime(6)                               not null,
    Version   timestamp(6) default current_timestamp(6) not null on update current_timestamp(6),
    IsDeleted tinyint(1)                                not null,
    DeletedAt datetime(6)                               null
);

create table `Order`.Orders
(
    Id                           bigint auto_increment
        primary key,
    SiteId                       bigint                                    not null,
    TotalRawPrice                bigint                                    not null,
    TotalCouponDiscount          bigint                                    not null,
    TotalPriceWithCouponDiscount bigint                                    not null,
    CourierPrice                 bigint                                    not null,
    Courier                      longtext                                  not null,
    OrderStatus                  longtext                                  not null,
    TotalFinalPrice              bigint                                    not null,
    Description                  longtext                                  null,
    TotalWeight                  int                                       not null,
    TrackingCode                 longtext                                  null,
    BasketId                     bigint                                    not null,
    DiscountId                   bigint                                    null,
    AddressId                    bigint                                    not null,
    CustomerId                   bigint                                    not null,
    CreatedAt                    datetime(6)                               not null,
    UpdatedAt                    datetime(6)                               not null,
    Version                      timestamp(6) default current_timestamp(6) not null on update current_timestamp(6),
    IsDeleted                    tinyint(1)                                not null,
    DeletedAt                    datetime(6)                               null
);

create table `Order`.OrderItems
(
    Id                           bigint auto_increment
        primary key,
    Quantity                     int                                       not null,
    RawPrice                     bigint                                    not null,
    FinalRawPrice                bigint                                    not null,
    FinalPriceWithCouponDiscount bigint                                    not null,
    JustCouponPrice              bigint                                    not null,
    JustDiscountPrice            bigint                                    not null,
    OrderId                      bigint                                    not null,
    ProductId                    bigint                                    not null,
    ProductVariantId             bigint                                    not null,
    CreatedAt                    datetime(6)                               not null,
    UpdatedAt                    datetime(6)                               not null,
    Version                      timestamp(6) default current_timestamp(6) not null on update current_timestamp(6),
    IsDeleted                    tinyint(1)                                not null,
    DeletedAt                    datetime(6)                               null,
    constraint FK_OrderItems_Orders_OrderId
        foreign key (OrderId) references `Order`.Orders (Id)
            on delete cascade
);

create index IX_OrderItems_OrderId
    on `Order`.OrderItems (OrderId);

create table Site.PageArticleUsages
(
    Id        bigint auto_increment
        primary key,
    PageId    bigint not null,
    ArticleId bigint not null,
    SiteId    bigint not null,
    UserId    bigint not null
);

create index IX_PageArticleUsages_ArticleId
    on Site.PageArticleUsages (ArticleId);

create index IX_PageArticleUsages_SiteId
    on Site.PageArticleUsages (SiteId);

create table Site.PageHeaderFooterUsages
(
    Id             bigint auto_increment
        primary key,
    PageId         bigint not null,
    HeaderFooterId bigint not null,
    SiteId         bigint not null,
    UserId         bigint not null
);

create index IX_PageHeaderFooterUsages_HeaderFooterId
    on Site.PageHeaderFooterUsages (HeaderFooterId);

create index IX_PageHeaderFooterUsages_SiteId
    on Site.PageHeaderFooterUsages (SiteId);

create table Site.PageProductUsages
(
    Id        bigint auto_increment
        primary key,
    PageId    bigint not null,
    ProductId bigint not null,
    SiteId    bigint not null,
    UserId    bigint not null
);

create index IX_PageProductUsages_ProductId
    on Site.PageProductUsages (ProductId);

create index IX_PageProductUsages_SiteId
    on Site.PageProductUsages (SiteId);

create table Site.Pages
(
    Id          bigint auto_increment
        primary key,
    SiteId      bigint                                    not null,
    HeaderId    bigint                                    not null,
    FooterId    bigint                                    not null,
    Slug        longtext                                  not null,
    Title       longtext                                  not null,
    Description longtext                                  null,
    Body        longtext collate utf8mb4_bin              null
        check (json_valid(`Body`)),
    SeoTags     longtext                                  null,
    UserId      bigint                                    not null,
    CreatedAt   datetime(6)                               not null,
    UpdatedAt   datetime(6)                               not null,
    Version     timestamp(6) default current_timestamp(6) not null on update current_timestamp(6),
    IsDeleted   tinyint(1)                                not null,
    DeletedAt   datetime(6)                               null
);

create table Site.PageMedia
(
    Id      bigint auto_increment
        primary key,
    PageId  bigint not null,
    MediaId bigint not null,
    constraint FK_PageMedia_Pages_PageId
        foreign key (PageId) references Site.Pages (Id)
            on delete cascade
);

create index IX_PageMedia_PageId
    on Site.PageMedia (PageId);

create table Payment.ParbadPayments
(
    Id                 bigint auto_increment
        primary key,
    TrackingNumber     bigint          not null,
    Amount             decimal(65, 30) not null,
    Token              longtext        null,
    TransactionCode    longtext        null,
    GatewayName        longtext        null,
    GatewayAccountName longtext        null,
    IsCompleted        tinyint(1)      not null,
    IsPaid             tinyint(1)      not null
);

create table Payment.ParbadTransactions
(
    Id             bigint auto_increment
        primary key,
    Amount         decimal(65, 30)  not null,
    Type           tinyint unsigned not null,
    IsSucceed      tinyint(1)       not null,
    Message        longtext         null,
    AdditionalData longtext         null,
    PaymentId      bigint           not null
);

create table Payment.Payments
(
    Id                  bigint auto_increment
        primary key,
    SiteId              bigint                                    not null,
    PaymentStatusEnum   longtext                                  not null,
    UserType            longtext                                  null,
    TrackingNumber      bigint                                    not null,
    Gateway             longtext                                  not null,
    GatewayAccountName  longtext                                  not null,
    Amount              bigint                                    not null,
    ServiceName         longtext                                  not null,
    ServiceAction       longtext                                  not null,
    OrderId             bigint                                    not null,
    ReturnUrl           longtext                                  not null,
    CallVerifyUrl       longtext                                  not null,
    ClientIp            longtext                                  not null,
    Message             longtext                                  null,
    GatewayResponseCode longtext                                  null,
    TransactionCode     longtext                                  null,
    AdditionalData      longtext                                  null,
    OrderData           longtext                                  null,
    UserId              bigint                                    not null,
    CustomerId          bigint                                    not null,
    CreatedAt           datetime(6)                               not null,
    UpdatedAt           datetime(6)                               not null,
    Version             timestamp(6) default current_timestamp(6) not null on update current_timestamp(6),
    IsDeleted           tinyint(1)                                not null,
    DeletedAt           datetime(6)                               null
);

create table User.Permissions
(
    Id   bigint auto_increment
        primary key,
    Name longtext not null
);

create table User.Plans
(
    Id               bigint auto_increment
        primary key,
    Name             longtext not null,
    ShowStatus       longtext not null,
    Description      longtext null,
    Price            bigint   not null,
    DiscountType     longtext null,
    Discount         bigint   null,
    Duration         int      not null,
    Feature          longtext null,
    SmsCredits       int      not null,
    EmailCredits     int      not null,
    StorageMbCredits int      not null,
    AiCredits        int      not null,
    AiImageCredits   int      not null
);

create table Product.ProductReviews
(
    Id         bigint auto_increment
        primary key,
    Rating     int                                       not null,
    `Like`     int                                       not null,
    Dislike    int                                       not null,
    Approved   tinyint(1)                                not null,
    ReviewText longtext                                  not null,
    ProductId  bigint                                    not null,
    SiteId     bigint                                    not null,
    UserId     bigint                                    not null,
    CustomerId bigint                                    not null,
    CreatedAt  datetime(6)                               not null,
    UpdatedAt  datetime(6)                               not null,
    Version    timestamp(6) default current_timestamp(6) not null on update current_timestamp(6),
    IsDeleted  tinyint(1)                                not null,
    DeletedAt  datetime(6)                               null
);

create table Product.Products
(
    Id              bigint auto_increment
        primary key,
    Name            longtext                                  not null,
    Description     longtext                                  null,
    Status          longtext                                  not null,
    Weight          int                                       not null,
    SellingCount    int                                       not null,
    VisitedCount    int                                       not null,
    ReviewCount     int                                       not null,
    Rate            int                                       not null,
    Badges          longtext                                  null,
    FreeSend        tinyint(1)                                not null,
    LongDescription longtext                                  null,
    Slug            longtext                                  not null,
    SeoTags         longtext                                  null,
    SiteId          bigint                                    not null,
    UserId          bigint                                    not null,
    CreatedAt       datetime(6)                               not null,
    UpdatedAt       datetime(6)                               not null,
    Version         timestamp(6) default current_timestamp(6) not null on update current_timestamp(6),
    IsDeleted       tinyint(1)                                not null,
    DeletedAt       datetime(6)                               null
);

create table Product.CategoryProduct
(
    Id         bigint auto_increment
        primary key,
    ProductId  bigint not null,
    CategoryId bigint not null,
    constraint FK_CategoryProduct_Categories_CategoryId
        foreign key (CategoryId) references Product.Categories (Id)
            on delete cascade,
    constraint FK_CategoryProduct_Products_ProductId
        foreign key (ProductId) references Product.Products (Id)
            on delete cascade
);

create index IX_CategoryProduct_CategoryId
    on Product.CategoryProduct (CategoryId);

create index IX_CategoryProduct_ProductId
    on Product.CategoryProduct (ProductId);

create table Product.Coupons
(
    Id         bigint auto_increment
        primary key,
    ProductId  bigint                                    not null,
    Quantity   int                                       not null,
    Type       longtext                                  not null,
    Value      bigint                                    not null,
    ExpiryDate datetime(6)                               not null,
    CreatedAt  datetime(6)                               not null,
    UpdatedAt  datetime(6)                               not null,
    Version    timestamp(6) default current_timestamp(6) not null on update current_timestamp(6),
    IsDeleted  tinyint(1)                                not null,
    DeletedAt  datetime(6)                               null,
    constraint IX_Coupons_ProductId
        unique (ProductId),
    constraint FK_Coupons_Products_ProductId
        foreign key (ProductId) references Product.Products (Id)
            on delete cascade
);

create table Product.DiscountProduct
(
    Id         bigint auto_increment
        primary key,
    ProductId  bigint not null,
    DiscountId bigint not null,
    constraint FK_DiscountProduct_Products_ProductId
        foreign key (ProductId) references Product.Products (Id)
            on delete cascade
);

create index IX_DiscountProduct_ProductId
    on Product.DiscountProduct (ProductId);

create table Product.ProductAttributes
(
    Id        bigint auto_increment
        primary key,
    ProductId bigint                                    not null,
    Type      longtext                                  not null,
    Name      longtext                                  not null,
    Value     longtext                                  not null,
    CreatedAt datetime(6)                               not null,
    UpdatedAt datetime(6)                               not null,
    Version   timestamp(6) default current_timestamp(6) not null on update current_timestamp(6),
    IsDeleted tinyint(1)                                not null,
    DeletedAt datetime(6)                               null,
    constraint FK_ProductAttributes_Products_ProductId
        foreign key (ProductId) references Product.Products (Id)
            on delete cascade
);

create index IX_ProductAttributes_ProductId
    on Product.ProductAttributes (ProductId);

create table Product.ProductMedia
(
    Id        bigint auto_increment
        primary key,
    ProductId bigint not null,
    MediaId   bigint not null,
    constraint FK_ProductMedia_Products_ProductId
        foreign key (ProductId) references Product.Products (Id)
            on delete cascade
);

create index IX_ProductMedia_ProductId
    on Product.ProductMedia (ProductId);

create table Product.ProductVariants
(
    Id         bigint auto_increment
        primary key,
    ProductId  bigint                                    not null,
    Name       longtext                                  not null,
    Price      bigint                                    not null,
    Stock      int                                       not null,
    UserId     bigint                                    not null,
    CustomerId bigint                                    not null,
    CreatedAt  datetime(6)                               not null,
    UpdatedAt  datetime(6)                               not null,
    Version    timestamp(6) default current_timestamp(6) not null on update current_timestamp(6),
    IsDeleted  tinyint(1)                                not null,
    DeletedAt  datetime(6)                               null,
    constraint FK_ProductVariants_Products_ProductId
        foreign key (ProductId) references Product.Products (Id)
            on delete cascade
);

create index IX_ProductVariants_ProductId
    on Product.ProductVariants (ProductId);

create table User.Provinces
(
    Id     bigint auto_increment
        primary key,
    Name   longtext not null,
    Slug   longtext not null,
    Status int      not null
);

create table User.Cities
(
    Id         bigint auto_increment
        primary key,
    Name       longtext                                  not null,
    Slug       longtext                                  not null,
    Status     longtext                                  not null,
    ProvinceId bigint                                    not null,
    Version    timestamp(6) default current_timestamp(6) not null on update current_timestamp(6),
    constraint FK_Cities_Provinces_ProvinceId
        foreign key (ProvinceId) references User.Provinces (Id)
            on delete cascade
);

create table User.Addresses
(
    Id          bigint auto_increment
        primary key,
    Title       longtext                                  null,
    Latitude    float                                     null,
    Longitude   float                                     null,
    AddressLine longtext                                  not null,
    PostalCode  longtext                                  not null,
    CityId      bigint                                    not null,
    ProvinceId  bigint                                    not null,
    UserId      bigint                                    not null,
    CustomerId  bigint                                    not null,
    CreatedAt   datetime(6)                               not null,
    UpdatedAt   datetime(6)                               not null,
    Version     timestamp(6) default current_timestamp(6) not null on update current_timestamp(6),
    IsDeleted   tinyint(1)                                not null,
    DeletedAt   datetime(6)                               null,
    constraint FK_Addresses_Cities_CityId
        foreign key (CityId) references User.Cities (Id),
    constraint FK_Addresses_Provinces_ProvinceId
        foreign key (ProvinceId) references User.Provinces (Id)
);

create table User.AddressCustomer
(
    Id         bigint auto_increment
        primary key,
    AddressId  bigint not null,
    CustomerId bigint not null,
    constraint FK_AddressCustomer_Addresses_AddressId
        foreign key (AddressId) references User.Addresses (Id)
            on delete cascade,
    constraint FK_AddressCustomer_Customers_CustomerId
        foreign key (CustomerId) references User.Customers (Id)
            on delete cascade
);

create index IX_AddressCustomer_AddressId
    on User.AddressCustomer (AddressId);

create index IX_AddressCustomer_CustomerId
    on User.AddressCustomer (CustomerId);

create index IX_Addresses_CityId
    on User.Addresses (CityId);

create index IX_Addresses_ProvinceId
    on User.Addresses (ProvinceId);

create index IX_Cities_ProvinceId
    on User.Cities (ProvinceId);

create table `Order`.ReturnItem
(
    Id           bigint auto_increment
        primary key,
    ReturnReason longtext                                  not null,
    OrderStatus  int                                       not null,
    OrderItemId  bigint                                    not null,
    ProductId    bigint                                    not null,
    UserId       bigint                                    not null,
    CustomerId   bigint                                    not null,
    CreatedAt    datetime(6)                               not null,
    UpdatedAt    datetime(6)                               not null,
    Version      timestamp(6) default current_timestamp(6) not null on update current_timestamp(6),
    IsDeleted    tinyint(1)                                not null,
    DeletedAt    datetime(6)                               null,
    constraint IX_ReturnItem_OrderItemId
        unique (OrderItemId),
    constraint FK_ReturnItem_OrderItems_OrderItemId
        foreign key (OrderItemId) references `Order`.OrderItems (Id)
            on delete cascade
);

create table User.Roles
(
    Id   bigint auto_increment
        primary key,
    Name longtext not null
);

create table User.CustomerRoles
(
    Id         bigint auto_increment
        primary key,
    RoleId     bigint not null,
    CustomerId bigint not null,
    constraint FK_CustomerRoles_Customers_CustomerId
        foreign key (CustomerId) references User.Customers (Id)
            on delete cascade,
    constraint FK_CustomerRoles_Roles_RoleId
        foreign key (RoleId) references User.Roles (Id)
            on delete cascade
);

create index IX_CustomerRoles_CustomerId
    on User.CustomerRoles (CustomerId);

create index IX_CustomerRoles_RoleId
    on User.CustomerRoles (RoleId);

create table User.PermissionRoles
(
    Id           bigint auto_increment
        primary key,
    RoleId       bigint not null,
    PermissionId bigint not null,
    constraint FK_PermissionRoles_Permissions_PermissionId
        foreign key (PermissionId) references User.Permissions (Id)
            on delete cascade,
    constraint FK_PermissionRoles_Roles_RoleId
        foreign key (RoleId) references User.Roles (Id)
            on delete cascade
);

create index IX_PermissionRoles_PermissionId
    on User.PermissionRoles (PermissionId);

create index IX_PermissionRoles_RoleId
    on User.PermissionRoles (RoleId);

create table User.RolePlan
(
    Id     bigint auto_increment
        primary key,
    RoleId bigint not null,
    PlanId bigint not null,
    constraint FK_RolePlan_Plans_PlanId
        foreign key (PlanId) references User.Plans (Id)
            on delete cascade,
    constraint FK_RolePlan_Roles_RoleId
        foreign key (RoleId) references User.Roles (Id)
            on delete cascade
);

create index IX_RolePlan_PlanId
    on User.RolePlan (PlanId);

create index IX_RolePlan_RoleId
    on User.RolePlan (RoleId);

create table Site.Sites
(
    Id         bigint auto_increment
        primary key,
    Domain     longtext                                  not null,
    DomainType longtext                                  not null,
    Name       longtext                                  not null,
    Status     longtext                                  not null,
    SiteType   longtext                                  not null,
    UserId     bigint                                    not null,
    CreatedAt  datetime(6)                               not null,
    UpdatedAt  datetime(6)                               not null,
    Version    timestamp(6) default current_timestamp(6) not null on update current_timestamp(6),
    IsDeleted  tinyint(1)                                not null,
    DeletedAt  datetime(6)                               null
);

create table Site.Settings
(
    Id         bigint auto_increment
        primary key,
    SiteId     bigint                                    not null,
    UserId     bigint                                    not null,
    CustomerId bigint                                    not null,
    CreatedAt  datetime(6)                               not null,
    UpdatedAt  datetime(6)                               not null,
    Version    timestamp(6) default current_timestamp(6) not null on update current_timestamp(6),
    IsDeleted  tinyint(1)                                not null,
    DeletedAt  datetime(6)                               null,
    constraint IX_Settings_SiteId
        unique (SiteId),
    constraint FK_Settings_Sites_SiteId
        foreign key (SiteId) references Site.Sites (Id)
            on delete cascade
);

create table Drive.Storages
(
    Id          bigint auto_increment
        primary key,
    UsedSpaceKb bigint                                    not null,
    QuotaKb     bigint                                    not null,
    ChargedAt   datetime(6)                               not null,
    ExpireAt    datetime(6)                               not null,
    UserId      bigint                                    not null,
    CreatedAt   datetime(6)                               not null,
    UpdatedAt   datetime(6)                               not null,
    Version     timestamp(6) default current_timestamp(6) not null on update current_timestamp(6),
    IsDeleted   tinyint(1)                                not null,
    DeletedAt   datetime(6)                               null
);

create table Support.Tickets
(
    Id         bigint auto_increment
        primary key,
    Title      longtext                                  not null,
    Status     longtext                                  not null,
    Category   longtext                                  not null,
    AssignedTo bigint                                    null,
    ClosedBy   bigint                                    null,
    ClosedAt   datetime(6)                               null,
    Priority   longtext                                  not null,
    UserId     bigint                                    not null,
    CreatedAt  datetime(6)                               not null,
    UpdatedAt  datetime(6)                               not null,
    Version    timestamp(6) default current_timestamp(6) not null on update current_timestamp(6),
    IsDeleted  tinyint(1)                                not null,
    DeletedAt  datetime(6)                               null
);

create table Support.Comments
(
    Id           bigint auto_increment
        primary key,
    TicketId     bigint                                    not null,
    Content      longtext                                  not null,
    RespondentId bigint                                    not null,
    CreatedAt    datetime(6)                               not null,
    UpdatedAt    datetime(6)                               not null,
    Version      timestamp(6) default current_timestamp(6) not null on update current_timestamp(6),
    IsDeleted    tinyint(1)                                not null,
    DeletedAt    datetime(6)                               null,
    constraint FK_Comments_Tickets_TicketId
        foreign key (TicketId) references Support.Tickets (Id)
            on delete cascade
);

create index IX_Comments_TicketId
    on Support.Comments (TicketId);

create table Support.TicketMedia
(
    Id       bigint auto_increment
        primary key,
    TicketId bigint not null,
    MediaId  bigint not null,
    constraint FK_TicketMedia_Tickets_TicketId
        foreign key (TicketId) references Support.Tickets (Id)
            on delete cascade
);

create index IX_TicketMedia_TicketId
    on Support.TicketMedia (TicketId);

create table User.UnitPrices
(
    Id           bigint auto_increment
        primary key,
    Name         longtext   not null,
    HasDay       tinyint(1) not null,
    Price        bigint     not null,
    DiscountType longtext   null,
    Discount     bigint     null
);

create table User.Users
(
    Id                       bigint auto_increment
        primary key,
    FirstName                longtext                                  null,
    LastName                 longtext                                  null,
    Email                    varchar(255)                              not null,
    AvatarId                 bigint                                    null,
    VerifyEmail              longtext                                  null,
    Password                 longtext                                  not null,
    Salt                     longtext                                  not null,
    NationalCode             longtext                                  null,
    Phone                    longtext                                  null,
    VerifyPhone              longtext                                  null,
    IsActive                 longtext                                  not null,
    AiTypeEnum               longtext                                  not null,
    UserTypeEnum             longtext                                  not null,
    PlanId                   bigint                                    null,
    PlanStartedAt            datetime(6)                               null,
    PlanExpiredAt            datetime(6)                               null,
    VerifyCode               int                                       null,
    ExpireVerifyCodeAt       datetime(6)                               null,
    AiCredits                int                                       not null,
    AiImageCredits           int                                       not null,
    StorageMbCredits         int                                       not null,
    StorageMbCreditsExpireAt datetime(6)                               null,
    EmailCredits             int                                       not null,
    SmsCredits               int                                       not null,
    UseCustomEmailSmtp       longtext                                  not null,
    Smtp_Host                longtext                                  null,
    Smtp_Port                int                                       null,
    Smtp_Username            longtext                                  null,
    Smtp_Password            longtext                                  null,
    Smtp_EnableSsl           tinyint(1)                                null,
    Smtp_SenderEmail         longtext                                  null,
    IsAdmin                  tinyint(1)                                not null,
    CreatedAt                datetime(6)                               not null,
    UpdatedAt                datetime(6)                               not null,
    Version                  timestamp(6) default current_timestamp(6) not null on update current_timestamp(6),
    IsDeleted                tinyint(1)                                not null,
    DeletedAt                datetime(6)                               null,
    constraint IX_Users_Email
        unique (Email)
);

create table User.AddressUser
(
    Id        bigint auto_increment
        primary key,
    AddressId bigint not null,
    UserId    bigint not null,
    constraint FK_AddressUser_Addresses_AddressId
        foreign key (AddressId) references User.Addresses (Id)
            on delete cascade,
    constraint FK_AddressUser_Users_UserId
        foreign key (UserId) references User.Users (Id)
            on delete cascade
);

create index IX_AddressUser_AddressId
    on User.AddressUser (AddressId);

create index IX_AddressUser_UserId
    on User.AddressUser (UserId);

create table User.RoleUser
(
    Id     bigint auto_increment
        primary key,
    RoleId bigint not null,
    UserId bigint not null,
    constraint FK_RoleUser_Roles_RoleId
        foreign key (RoleId) references User.Roles (Id)
            on delete cascade,
    constraint FK_RoleUser_Users_UserId
        foreign key (UserId) references User.Users (Id)
            on delete cascade
);

create index IX_RoleUser_RoleId
    on User.RoleUser (RoleId);

create index IX_RoleUser_UserId
    on User.RoleUser (UserId);

